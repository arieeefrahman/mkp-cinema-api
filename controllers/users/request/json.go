package request

import (
	"errors"
	"fmt"
	"mkp-cinema-api/businesses/users"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Username             string `json:"username" validate:"required,lowercase"`
	Password             string `json:"password" validate:"required"`
	ConfirmationPassword string `json:"confirmation_password" validate:"required"`
	Email                string `json:"email" validate:"required,email,lowercase"`
}

type UserLogin struct {
	Username string `json:"username" validate:"required,lowercase"`
	Password string `json:"password" validate:"required"`
}

func (req *User) ToDomainRegister() *users.Domain {
	return &users.Domain{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *UserLogin) ToDomainLogin() *users.LoginDomain {
	return &users.LoginDomain{
		Username: req.Username,
		Password: req.Password,
	}
}

func (req *User) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return formatValidationError(err)
	}
	return nil
}

func (req *UserLogin) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return formatValidationError(err)
	}
	return nil
}

func formatValidationError(err error) error {
	if errs, ok := err.(validator.ValidationErrors); ok {
		var errorMsg string
		for _, e := range errs {
			switch e.Tag() {
			case "required":
				errorMsg += fmt.Sprintf("Field '%s' is required.\n", e.Field())
			case "email":
				errorMsg += fmt.Sprintf("Field '%s' must be a valid email address.\n", e.Field())
			case "lowercase":
				errorMsg += fmt.Sprintf("Field '%s' must be in lowercase.\n", e.Field())
			default:
				errorMsg += fmt.Sprintf("Field '%s' is invalid.\n", e.Field())
			}
		}
		return errors.New(errorMsg)
	}
	
	return err
}
