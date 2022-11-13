package repository

import (
	"context"
	"test/cmd/entity"
	"test/cmd/model"
)

type UserRepository interface {
	Insert(ctx context.Context, user *entity.User) error
	FindAll(ctx context.Context, param model.CommonParam) ([]entity.User, error)
	FindByID(ctx context.Context, id int64) (*entity.User, error)
	Count(ctx context.Context, param model.CommonParam) (int64, error)
}
