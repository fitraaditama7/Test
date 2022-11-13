package mapper

import (
	"test/cmd/entity"
	"test/cmd/model"
	"time"
)

func CreateUserRequestToUserEntity(request model.RegisterUserRequest) *entity.User {
	return &entity.User{
		Name:              request.Name,
		Address:           request.Address,
		Email:             request.Email,
		Password:          request.Password,
		Photos:            request.Photos,
		CreditCardType:    request.CreditCardType,
		CreditCardNumber:  request.CreditCardNumber,
		CreditCardName:    request.CreditCardName,
		CreditCardExpired: request.CreditCardExpired,
		CreditCardCVV:     request.CreditCardCVV,
		CreatedAt:         time.Now(),
	}
}

func UserEntityToCreateUserResponse(user *entity.User) *model.RegisterUserResponse {
	return &model.RegisterUserResponse{UserID: user.ID}
}

func UserEntityToModelUser(user entity.User) *model.User {
	return &model.User{
		UserID:  user.ID,
		Name:    user.Name,
		Address: user.Address,
		Email:   user.Email,
		Photos:  user.Photos,
		CreditCard: model.UserCreditCard{
			Type:    user.CreditCardType,
			Number:  user.CreditCardNumber,
			Name:    user.Name,
			Expired: user.CreditCardExpired,
			CVV:     user.CreditCardCVV,
		},
	}
}

func ModelUsersToModelListUser(count int64, users []*model.User) *model.ListUser {
	return &model.ListUser{
		Count: count,
		Rows:  users,
	}
}
