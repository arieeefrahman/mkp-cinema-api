package studios

import (
	"mkp-cinema-api/businesses/studios"
	"mkp-cinema-api/drivers/postgresql/cinemas"
	"time"

	"gorm.io/gorm"
)

type Studio struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
	TotalSeat uint
	CinemaID  uint           `gorm:"not null"`
	Cinema    cinemas.Cinema `gorm:"foreignKey:CinemaID"`
}

func FromDomain(domain *studios.Domain) *Studio {
	return &Studio{
		ID:        domain.ID,
		Name:      domain.Name,
		TotalSeat: domain.TotalSeat,
		CinemaID:  domain.CinemaID,
	}
}

func (rec *Studio) ToDomain() studios.Domain {
	return studios.Domain{
		ID:        rec.ID,
		Name:      rec.Name,
		TotalSeat: rec.TotalSeat,
		CinemaID:  rec.CinemaID,
	}
}
