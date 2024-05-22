package response

import (
	"mkp-cinema-api/businesses/movies"
	"time"
)

type Movie struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title" form:"title"`
	Genre       string    `json:"genre" form:"genre"`
	Duration    uint      `json:"duration" form:"duration"`
	Rating      float32   `json:"rating" form:"rating"`
	ReleaseDate time.Time `json:"release_date" form:"release_date"`
	Description string    `json:"description" form:"description"`
}

func FromDomain(domain movies.Domain) Movie {
	return Movie{
		ID:          domain.ID,
		Title:       domain.Title,
		Genre:       domain.Genre,
		Duration:    domain.Duration,
		Rating:      domain.Rating,
		ReleaseDate: domain.ReleaseDate,
		Description: domain.Description,
	}
}
