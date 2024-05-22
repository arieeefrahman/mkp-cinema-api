package cinemas

import "time"

type Domain struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	Location  string
	CityID    uint
}

type Usecase interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(cinemaDomain *Domain) (Domain, error)
	Update(id string, cinemaDomain *Domain) (Domain, error)
	Delete(id string) (bool, error)
}

type Repository interface {
	GetAll() ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(cinemaDomain *Domain) (Domain, error)
	Update(id string, cinemaDomain *Domain) (Domain, error)
	Delete(id string) (bool, error)
}
