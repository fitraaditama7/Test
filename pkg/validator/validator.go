package validator

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"log"
	"reflect"
	"strings"
)

var validate = validator.New()

func validateMessage(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return fe.Field() + " is required"
	case "email":
		return "Invalid email"
	}
	return fe.Error() // default error
}

func Validate(a interface{}) error {
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	err := validate.Struct(a)
	if err != nil {
		log.Println(err)
		validationErr := err.(validator.ValidationErrors)[0]
		return errors.New(validateMessage(validationErr))
	}
	return nil
}
