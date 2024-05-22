package cities

import (
	"mkp-cinema-api/businesses/cities"
	"time"
)

type City struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
}

func FromDomain(domain *cities.Domain) *City {
	return &City{
		ID:   domain.ID,
		Name: domain.Name,
	}
}

func (rec *City) ToDomain() cities.Domain {
	return cities.Domain{
		ID:   rec.ID,
		Name: rec.Name,
	}
}
