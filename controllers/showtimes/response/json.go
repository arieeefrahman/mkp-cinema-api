package response

import (
	"mkp-cinema-api/businesses/showtimes"
	"time"
)

type Showtime struct {
	ID        uint      `json:"id"`
	Date      time.Time `json:"date" form:"date" validate:"required"`
	StartTime string    `json:"start_time" form:"start_time" validate:"required"`
	EndTime   string    `json:"end_time" form:"end_time" validate:"required"`
	MovieID   uint      `json:"movie_id" form:"movie_id" validate:"required"`
	StudioID  uint      `json:"studio_id" form:"studio_id" validate:"required"`
	CinemaID  uint      `json:"cinema_id" form:"cinema_id" validate:"required"`
}

func FromDomain(domain showtimes.Domain) Showtime {
	return Showtime{
		ID:        domain.ID,
		Date:      domain.Date,
		StartTime: domain.StartTime.Format("15:04"),
		EndTime:   domain.EndTime.Format("15:04"),
		MovieID:   domain.MovieID,
		StudioID:  domain.StudioID,
		CinemaID:  domain.CinemaID,
	}
}
