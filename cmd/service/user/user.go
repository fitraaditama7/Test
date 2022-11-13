package userservice

import (
	"context"
	"log"
	"test/cmd/constant"
	"test/cmd/mapper"
	"test/cmd/model"
	"test/cmd/repository"
	"test/cmd/service"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) service.UserService {
	return &userService{userRepository: userRepository}
}

func (u *userService) Register(ctx context.Context, user model.RegisterUserRequest) (*model.RegisterUserResponse, error) {
	userEntity := mapper.CreateUserRequestToUserEntity(user)
	err := u.userRepository.Insert(ctx, userEntity)
	if err != nil {
		log.Println(err)
		return nil, constant.SystemError
	}

	return mapper.UserEntityToCreateUserResponse(userEntity), nil
}

func (u *userService) List(ctx context.Context, param model.CommonParam) (*model.ListUser, error) {
	countUser, err := u.userRepository.Count(ctx, param)
	if err != nil {
		log.Println(err)
		return nil, constant.SystemError
	}

	users, err := u.userRepository.FindAll(ctx, param)
	if err != nil {
		log.Println(err)
		return nil, constant.SystemError
	}

	var modelUsers []*model.User
	for _, user := range users {
		modelUser := mapper.UserEntityToModelUser(user)
		modelUsers = append(modelUsers, modelUser)
	}

	if len(modelUsers) == 0 {
		modelUsers = make([]*model.User, 0)
	}

	return mapper.ModelUsersToModelListUser(countUser, modelUsers), nil
}

func (u *userService) Detail(ctx context.Context, id int64) (*model.User, error) {
	user, err := u.userRepository.FindByID(ctx, id)
	if err != nil {
		log.Println(err)
		return nil, constant.SystemError
	}

	if user == nil {
		return nil, constant.UserNotFound
	}

	modelUser := mapper.UserEntityToModelUser(*user)

	return modelUser, nil
}
