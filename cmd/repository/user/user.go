package userrepository

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"test/cmd/entity"
	"test/cmd/model"
	"test/cmd/repository"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (u userRepository) Insert(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Create(&user).Error
}

func (u userRepository) FindAll(ctx context.Context, param model.CommonParam) ([]entity.User, error) {
	var users []entity.User
	query := u.db

	if param.QueryBy != "" {
		query = query.Where("name LIKE ?", "%"+param.QueryBy+"%")
	}

	result := query.WithContext(ctx).Order(fmt.Sprintf("%s %s", param.OrderBy, param.SortBy)).Limit(param.Limit).Offset(param.Offset).Find(&users)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, result.Error
	}

	return users, nil
}

func (u userRepository) FindByID(ctx context.Context, id int64) (*entity.User, error) {
	var user *entity.User

	result := u.db.WithContext(ctx).Where("id = ?", id).First(&user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, result.Error
	}

	return user, nil
}

func (u userRepository) Count(ctx context.Context, param model.CommonParam) (int64, error) {
	var total int64
	query := u.db.Table("users")

	if param.QueryBy != "" {
		query = query.Where("name LIKE ?", "%"+param.QueryBy+"%")
	}

	result := query.WithContext(ctx).Count(&total)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return 0, nil
		}

		return 0, result.Error
	}

	return total, nil
}
