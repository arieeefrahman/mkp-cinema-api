package showtimes

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Date      time.Time
	StartTime time.Time
	EndTime   time.Time
	MovieID   uint `gorm:"foreignKey:MovieID"`
	StudioID  uint `gorm:"foreignKey:StudioID"`
	CinemaID  uint `gorm:"foreignKey:CinemaID"`
}

type Usecase interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(showtimeDomain *Domain) (Domain, error)
	Update(id string, showtimeDomain *Domain) (Domain, error)
	Delete(id string) (bool, error)
}

type Repository interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(showtimeDomain *Domain) (Domain, error)
	Update(id string, showtimeDomain *Domain) (Domain, error)
	Delete(id string) (bool, error)
}
