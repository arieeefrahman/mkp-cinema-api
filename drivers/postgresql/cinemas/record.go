package cinemas

import (
	"mkp-cinema-api/businesses/cinemas"
	"mkp-cinema-api/drivers/postgresql/cities"
	"time"
)

type Cinema struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Location  string
	CityID    uint        `gorm:"not null"`
	City      cities.City `gorm:"foreignKey:CityID"`
}

func FromDomain(domain *cinemas.Domain) *Cinema {
	return &Cinema{
		ID:       domain.ID,
		Name:     domain.Name,
		Location: domain.Location,
		CityID:   domain.CityID,
	}
}

func (rec *Cinema) ToDomain() cinemas.Domain {
	return cinemas.Domain{
		ID:   rec.ID,
		Name: rec.Name,
		Location: rec.Location,
		CityID: rec.CityID,
	}
}
