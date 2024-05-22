package request

import (
	"errors"
	"fmt"
	"mkp-cinema-api/businesses/cities"

	"github.com/go-playground/validator/v10"
)

type CityCreate struct {
	Name string `json:"name" form:"name" validate:"required"`
}

type CityUpdate struct {
	Name *string `json:"name" form:"name" validate:"omitempty"`
}

func (req *CityCreate) ToDomain() *cities.Domain {
	return &cities.Domain{
		Name: req.Name,
	}
}

func (req *CityUpdate) ToDomain() *cities.Domain {
	domain := &cities.Domain{}
	if req.Name != nil {
		domain.Name = *req.Name
	}

	return domain
}

func (req *CityCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return formatValidationError(err)
	}
	return nil
}

func (req *CityUpdate) Validate() error {
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
			default:
				errorMsg += fmt.Sprintf("Field '%s' is invalid.\n", e.Field())
			}
		}
		return errors.New(errorMsg)
	}

	return err
}
