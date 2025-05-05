package user

import "github.com/jeff3710/ndot/internal/apiserver/service"

type UserHandler struct {
	service service.UserServiceInterface
}

func NewUserHandler(service service.UserServiceInterface) *UserHandler {
	return &UserHandler{
		service: service,
	}
}
