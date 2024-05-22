package cities

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Name      string
}

type Usecase interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(cityDomain *Domain) (Domain, error)
	Update(id string, cityDomain *Domain) (Domain, error)
	Delete(id string) (bool, error)
}

type Repository interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(cityDomain *Domain) (Domain, error)
	Update(id string, cityDomain *Domain) (Domain, error)
	Delete(id string) (bool, error)
}
