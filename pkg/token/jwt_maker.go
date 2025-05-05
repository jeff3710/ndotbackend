package token

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

const minSecretKeySize = 32

type JWTMaker struct {
	secretKey string
}

func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSecretKeySize {
		return nil, fmt.Errorf("无效的密钥长度，必须至少为 %d 个字符", minSecretKeySize)
	}
	return &JWTMaker{secretKey}, nil
}

// CreateToken implements Maker.
func (maker *JWTMaker) CreateToken(username string, role string, duration time.Duration,tokenType TokenType) (string, *Payload, error) {
	payload, err := NewPayload(username, role, duration,tokenType)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(maker.secretKey))
	return token, payload, err
}

// VerifyToken implements Maker.
func (maker *JWTMaker) VerifyToken(token string,tokenType TokenType) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}
	payload, ok := jwtToken.Claims.(*Payload)

	if !ok {
		return nil, ErrInvalidToken
	}
	err=payload.Valid(tokenType)
	if err!= nil {
		return nil, err
	}
	return payload, nil
}
