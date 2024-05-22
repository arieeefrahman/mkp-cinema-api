package movies

import (
	"mkp-cinema-api/businesses/movies"
	"time"

	"gorm.io/gorm"
)

type Movie struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
	Title       string
	Genre       string
	Duration    uint
	Rating      float32
	ReleaseDate time.Time
	Description string `gorm:"type:text"`
}

func FromDomain(domain *movies.Domain) *Movie {
	return &Movie{
		ID:          domain.ID,
		Title:       domain.Title,
		Genre:       domain.Genre,
		Duration:    domain.Duration,
		Rating:      domain.Rating,
		ReleaseDate: domain.ReleaseDate,
		Description: domain.Description,
	}
}

func (rec *Movie) ToDomain() movies.Domain {
	return movies.Domain{
		ID:          rec.ID,
		Title:       rec.Title,
		Genre:       rec.Genre,
		Duration:    rec.Duration,
		Rating:      rec.Rating,
		ReleaseDate: rec.ReleaseDate,
		Description: rec.Description,
	}
}
