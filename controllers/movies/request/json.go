package request

import (
	"errors"
	"fmt"
	"mkp-cinema-api/businesses/movies"
	"time"

	"github.com/go-playground/validator/v10"
)

type MovieCreate struct {
	Title       string    `json:"title" form:"title" validate:"required"`
	Genre       string    `json:"genre" form:"genre" validate:"required"`
	Duration    uint      `json:"duration" form:"duration" validate:"required"`
	Rating      float32   `json:"rating" form:"rating" validate:"required"`
	ReleaseDate time.Time `json:"release_date" form:"release_date" validate:"required"`
	Description string    `json:"description" form:"description" validate:"required"`
}

type MovieUpdate struct {
	Title       *string    `json:"title" form:"title" validate:"omitempty"`
	Genre       *string    `json:"genre" form:"genre" validate:"omitempty"`
	Duration    *uint      `json:"duration" form:"duration" validate:"omitempty"`
	Rating      *float32   `json:"rating" form:"rating" validate:"omitempty"`
	ReleaseDate *time.Time `json:"release_date" form:"release_date" validate:"omitempty"`
	Description *string    `json:"description" form:"description" validate:"omitempty"`
}

func (req *MovieCreate) ToDomain() *movies.Domain {
	return &movies.Domain{
		Title:       req.Title,
		Genre:       req.Genre,
		Duration:    req.Duration,
		Rating:      req.Rating,
		ReleaseDate: req.ReleaseDate,
		Description: req.Description,
	}
}

func (req *MovieUpdate) ToDomain() *movies.Domain {
	domain := &movies.Domain{}
	if req.Title != nil {
		domain.Title = *req.Title
	}
	if req.Genre != nil {
		domain.Genre = *req.Genre
	}
	if req.Duration != nil {
		domain.Duration = *req.Duration
	}
	if req.Rating != nil {
		domain.Rating = *req.Rating
	}
	if req.ReleaseDate != nil {
		domain.ReleaseDate = *req.ReleaseDate
	}
	if req.Description != nil {
		domain.Description = *req.Description
	}
	return domain
}

func (req *MovieCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return formatValidationError(err)
	}
	return nil
}

func (req *MovieUpdate) Validate() error {
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
