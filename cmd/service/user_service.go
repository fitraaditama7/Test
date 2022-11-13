package service

import (
	"context"
	"test/cmd/model"
)

type UserService interface {
	Register(ctx context.Context, user model.RegisterUserRequest) (*model.RegisterUserResponse, error)
	List(ctx context.Context, param model.CommonParam) (*model.ListUser, error)
	Detail(ctx context.Context, id int64) (*model.User, error)
}
