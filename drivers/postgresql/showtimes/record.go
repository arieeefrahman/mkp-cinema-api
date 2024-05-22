package showtimes

import (
	"mkp-cinema-api/businesses/showtimes"
	"mkp-cinema-api/drivers/postgresql/cinemas"
	"mkp-cinema-api/drivers/postgresql/movies"
	"mkp-cinema-api/drivers/postgresql/studios"
	"time"

	"gorm.io/gorm"
)

type Showtime struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Date      time.Time
	StartTime time.Time
	EndTime   time.Time
	MovieID   uint           `gorm:"foreignKey:MovieID"`
	StudioID  uint           `gorm:"foreignKey:StudioID"`
	CinemaID  uint           `gorm:"foreignKey:CinemaID"`
	Movie     movies.Movie   `gorm:"foreignKey:StudioID"`
	Studio    studios.Studio `gorm:"foreignKey:StudioID"`
	Cinema    cinemas.Cinema `gorm:"foreignKey:CinemaID"`
}

func FromDomain(domain *showtimes.Domain) *Showtime {
	return &Showtime{
		ID:        domain.ID,
		Date:      domain.Date,
		StartTime: domain.StartTime,
		EndTime:   domain.EndTime,
		MovieID:   domain.MovieID,
		StudioID:  domain.StudioID,
		CinemaID:  domain.CinemaID,
	}
}

func (rec *Showtime) ToDomain() showtimes.Domain {
	return showtimes.Domain{
		ID:        rec.ID,
		Date:      rec.Date,
		StartTime: rec.StartTime,
		EndTime:   rec.EndTime,
		MovieID:   rec.MovieID,
		StudioID:  rec.StudioID,
		CinemaID:  rec.CinemaID,
	}
}
