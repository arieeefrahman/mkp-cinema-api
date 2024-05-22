package movies

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID          uint
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

type Usecase interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(movieDomain *Domain) (Domain, error)
	Update(id string, movieDomain *Domain) (Domain, error)
	Delete(id string) (bool, error)
}

type Repository interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(movieDomain *Domain) (Domain, error)
	Update(id string, movieDomain *Domain) (Domain, error)
	Delete(id string) (bool, error)
}
