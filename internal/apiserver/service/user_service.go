package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	db "github.com/jeff3710/ndot/db/sqlc"
	"github.com/jeff3710/ndot/internal/pkg/core"
	"github.com/jeff3710/ndot/internal/pkg/errno"
	"github.com/jeff3710/ndot/pkg/config"
	"github.com/jeff3710/ndot/pkg/log"
	"github.com/jeff3710/ndot/pkg/token"
	"github.com/jeff3710/ndot/util"
)

type UserService struct {
	db     db.Store
	maker  token.Maker
	config *config.Config
}

var _ UserServiceInterface = (*UserService)(nil)

func NewUserService(db db.Store, maker token.Maker, config *config.Config) *UserService {
	return &UserService{
		db:     db,
		maker:  maker,
		config: config,
	}
}

func (s *UserService) CreateUser(ctx *gin.Context, req db.CreateUserParams) error {
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		log.Errorw("failed to hash password", "err", err.Error())
		core.WriteResponse(ctx, err, nil)
	}
	params := db.CreateUserParams{
		Username: req.Username,
		Password: hashedPassword,
		Role:     req.Role,
		Active:   req.Active,
	}
	if err := s.db.CreateUser(ctx, params); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil

}

type userResponse struct {
	Id                int32     `json:"id"`
	Username          string    `json:"username"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		Id:        user.ID,
		Username:  user.Username,
		CreatedAt: user.CreatedAt,
	}
}

type LoginUserResponse struct {
	SessionID             uuid.UUID    `json:"session_id"`
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiresAt  time.Time    `json:"access_token_expires_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time    `json:"refresh_token_expires_at"`
	User                  userResponse `json:"user"`
}

func (s *UserService) Login(ctx *gin.Context, username, password string) (*LoginUserResponse, error) {
	user, err := s.db.GetUser(ctx, username)
	if err != nil {

		return nil, errno.ErrUserNotFound
	}
	if err = util.CheckPassword(password, user.Password); err != nil {
		return nil, errno.ErrPasswordIncorrect
	}

	accessToken, accessPayload, err := s.maker.CreateToken(
		user.ID,
		user.Username,
		user.Role,
		s.config.Jwt.AccessTokenDuration,
		token.TokenTypeAccessToken,
	)
	if err != nil {
		return nil, errno.InternalServerError
	}

	refreshToken, refreshPayload, err := s.maker.CreateToken(
		user.ID,
		user.Username,
		user.Role,
		s.config.Jwt.RefreshTokenDuration,
		token.TokenTypeRefreshToken,
	)
	if err != nil {
		return nil, errno.InternalServerError
	}

	session, err := s.db.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    ctx.Request.UserAgent(),
		ClientIp:     ctx.ClientIP(),
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})
	if err != nil {
		core.WriteResponse(ctx, errno.InternalServerError, nil)
	}
	resp := LoginUserResponse{
		SessionID:             session.ID,
		AccessToken:           accessToken,
		AccessTokenExpiresAt:  accessPayload.ExpiredAt,
		RefreshToken:          refreshToken,
		RefreshTokenExpiresAt: refreshPayload.ExpiredAt,
		User:                  newUserResponse(user),
	}

	return &resp, nil
}

type GetUserResponse struct {
	Id        int32     `json:"id"`
	Username  string    `json:"username"`
	Role      string    `json:"role"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func newGetUserResponse(user db.User) GetUserResponse {
	return GetUserResponse{
		Id:        user.ID,
		Username:  user.Username,
		Role:      user.Role,
		Active:    user.Active,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (s *UserService) GetUser(ctx *gin.Context, username string) (*GetUserResponse, error) {
	user, err := s.db.GetUser(ctx, username)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}
	resp := newGetUserResponse(user)
	return &resp, nil

}
func (s *UserService) GetUserById(ctx *gin.Context, userId string) (*GetUserResponse, error) {
	id, err := util.StringToInt32(userId)
	if err != nil {
		return nil, fmt.Errorf("invalid user id: %w", err) // 或者其他错误处理逻辑，如返回 400 Bad Reques
	}
	user, err := s.db.GetUserById(ctx, id)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}
	resp := newGetUserResponse(user)
	return &resp, nil
}

type ListUserResponse struct {
	Users []GetUserResponse `json:"users"`
}

func (s *UserService) ListUser(ctx *gin.Context) (*ListUserResponse, error) {
	users, err := s.db.GetAllUsers(ctx)
	if err != nil {
		return nil, errno.ErrUserNotFound
	}
	resp := ListUserResponse{
		Users: make([]GetUserResponse, len(users)),
	}
	for i, user := range users {
		resp.Users[i] = newGetUserResponse(user)
	}
	return &resp, nil
}

func (s *UserService) DeleteUser(ctx *gin.Context, username string) error {
	if username == "admin" {
		return errno.ErrForbiddenDeleteAdmin
	}
	if err := s.db.DeleteUser(ctx, username); err != nil {
		return errno.ErrUserNotFound
	}
	return nil
}

func (s *UserService) RenewAccessToken(ctx *gin.Context, refreshToken string) (string, *token.Payload, error) {
	refreshPayload, err := s.maker.VerifyToken(refreshToken, token.TokenTypeRefreshToken)
	if err != nil {
		return "", nil, errno.ErrTokenInvalid
	}

	session, err := s.db.GetSession(ctx, refreshPayload.ID)
	if err != nil {
		return "", nil, err
	}
	if session.IsBlocked {
		return "", nil, fmt.Errorf("session is blocked")
	}
	if session.Username != refreshPayload.Username {
		return "", nil, fmt.Errorf("incorrect session user")
	}
	if session.RefreshToken != refreshToken {
		return "", nil, fmt.Errorf("mismatched session token")
	}
	if time.Now().After(session.ExpiresAt) {
		return "", nil, fmt.Errorf("expired session")
	}
	accessToken, accessPayload, err := s.maker.CreateToken(
		refreshPayload.UserID,
		refreshPayload.Username,
		refreshPayload.Role,
		s.config.Jwt.AccessTokenDuration,
		token.TokenTypeAccessToken,
	)
	if err != nil {
		return "", nil, err
	}
	return accessToken, accessPayload, nil
}
