package request

import (
	"errors"
	"fmt"
	"mkp-cinema-api/businesses/showtimes"
	"time"

	"github.com/go-playground/validator/v10"
)

// Custom Date type
type Date time.Time

const dateFormat = "2006-01-02"

func (d *Date) UnmarshalJSON(b []byte) error {
	strInput := string(b)
	parsedTime, err := time.Parse(`"`+dateFormat+`"`, strInput)
	if err != nil {
		return err
	}
	*d = Date(parsedTime)
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d Date) String() string {
	t := time.Time(d)
	return fmt.Sprintf(`"%s"`, t.Format(dateFormat))
}

// Custom Time type
type Time time.Time

const timeFormat = "15:04"

func (t *Time) UnmarshalJSON(b []byte) error {
	strInput := string(b)
	parsedTime, err := time.Parse(`"`+timeFormat+`"`, strInput)
	if err != nil {
		return err
	}
	*t = Time(parsedTime)
	return nil
}

func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(t.String()), nil
}

func (t Time) String() string {
	tm := time.Time(t)
	return fmt.Sprintf(`"%s"`, tm.Format(timeFormat))
}

type ShowtimeCreate struct {
	Date      Date  `json:"date" form:"date" validate:"required"`
	StartTime Time  `json:"start_time" form:"start_time" validate:"required"`
	EndTime   Time  `json:"end_time" form:"end_time" validate:"required"`
	MovieID   uint  `json:"movie_id" form:"movie_id" validate:"required"`
	StudioID  uint  `json:"studio_id" form:"studio_id" validate:"required"`
	CinemaID  uint  `json:"cinema_id" form:"cinema_id" validate:"required"`
}

type ShowtimeUpdate struct {
	Date      *Date  `json:"date" form:"date" validate:"omitempty"`
	StartTime *Time  `json:"start_time" form:"start_time" validate:"omitempty"`
	EndTime   *Time  `json:"end_time" form:"end_time" validate:"omitempty"`
	MovieID   *uint  `json:"movie_id" form:"movie_id" validate:"omitempty"`
	StudioID  *uint  `json:"studio_id" form:"studio_id" validate:"omitempty"`
	CinemaID  *uint  `json:"cinema_id" form:"cinema_id" validate:"omitempty"`
}

func (req *ShowtimeCreate) ToDomain() *showtimes.Domain {
	return &showtimes.Domain{
		Date:      time.Time(req.Date),
		StartTime: time.Time(req.StartTime),
		EndTime:   time.Time(req.EndTime),
		MovieID:   req.MovieID,
		StudioID:  req.StudioID,
		CinemaID:  req.CinemaID,
	}
}

func (req *ShowtimeUpdate) ToDomain() *showtimes.Domain {
	domain := &showtimes.Domain{}
	
	if req.Date != nil {
		domain.Date = time.Time(*req.Date)
	}
	if req.StartTime != nil {
		domain.StartTime = time.Time(*req.StartTime)
	}
	if req.EndTime != nil {
		domain.EndTime = time.Time(*req.EndTime)
	}
	
	return domain
}

func (req *ShowtimeCreate) Validate() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		return formatValidationError(err)
	}
	return nil
}

func (req *ShowtimeUpdate) Validate() error {
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
