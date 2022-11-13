package model

import (
	"log"
	"strconv"
	"strings"
	"test/cmd/constant"
	"test/pkg/utils"
	"test/pkg/validator"
)

type User struct {
	UserID     uint           `json:"user_id"`
	Name       string         `json:"name"`
	Address    string         `json:"address"`
	Email      string         `json:"email"`
	Photos     []string       `json:"photos"`
	CreditCard UserCreditCard `json:"credit_card"`
}

type UserCreditCard struct {
	Type    string `json:"type"`
	Number  string `json:"number"`
	Name    string `json:"name"`
	Expired string `json:"expired"`
	CVV     string `json:"cvv"`
}

type ListUser struct {
	Count int64   `json:"count"`
	Rows  []*User `json:"rows"`
}

type RegisterUserRequest struct {
	Name              string   `json:"name" validate:"required"`
	Address           string   `json:"address" validate:"required"`
	Email             string   `json:"email" validate:"required"`
	Password          string   `json:"password" validate:"required"`
	Photos            []string `json:"photos" validate:"required"`
	CreditCardType    string   `json:"credit_card_type" validate:"required"`
	CreditCardNumber  string   `json:"credit_card_number" validate:"required"`
	CreditCardName    string   `json:"credit_card_name" validate:"required"`
	CreditCardExpired string   `json:"credit_card_expired" validate:"required"`
	CreditCardCVV     string   `json:"credit_card_cvv" validate:"required"`
}

func (c *RegisterUserRequest) Validate() error {
	err := validator.Validate(c)
	if err != nil {
		log.Println(err)
		return constant.CustomBadRequest(err)
	}

	if !utils.IsEmailValid(c.Email) {
		return constant.InvalidEmail
	}

	if !strings.Contains(c.CreditCardNumber, constant.CreditCardSeparator) {
		return constant.InvalidCreditCardFormat
	}

	exp := strings.Split(c.CreditCardExpired, "/")
	expMonth, err := strconv.Atoi(exp[0])
	if err != nil {
		log.Println(err)
		return constant.InvalidCreditCardExpireFormat
	}
	expYear, err := strconv.Atoi(exp[1])
	if err != nil {
		log.Println(err)
		return constant.InvalidCreditCardExpireFormat
	}

	if expMonth < 1 || 12 < expMonth {
		return constant.InvalidCreditCardData
	}

	if expYear < 1900 || expYear > 2200 {
		return constant.InvalidCreditCardData
	}
	return nil
}

type RegisterUserResponse struct {
	UserID uint `json:"user_id"`
}
