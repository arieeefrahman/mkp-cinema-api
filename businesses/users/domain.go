package users

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Domain struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Username  string
	Password  string
	Email     string
}

type LoginDomain struct {
	Username string
	Password string
}

type Usecase interface {
	Register(userDomain *Domain) (Domain, error)
	Login(userDomain *LoginDomain) (map[string]string, error)
}

type Repository interface {
	Register(userDomain *Domain) (Domain, error)
	GetByUsername(userDomain *LoginDomain) (Domain, error)
}