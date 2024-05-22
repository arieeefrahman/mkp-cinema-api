package response

import (
	"mkp-cinema-api/businesses/cities"
)

type City struct {
	ID   uint   `json:"id"`
	Name string `json:"Name" form:"Name"`
}

func FromDomain(domain cities.Domain) City {
	return City{
		ID:   domain.ID,
		Name: domain.Name,
	}
}
