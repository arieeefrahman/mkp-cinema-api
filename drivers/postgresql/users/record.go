package users

import (
	"mkp-cinema-api/businesses/users"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	Username  string         `json:"username"`
	Password  string         `json:"password"`
	Email     string         `json:"email"`
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Username:  domain.Username,
		Password:  domain.Password,
		Email:     domain.Email,
	}
}

func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		Username:  rec.Username,
		Password:  rec.Password,
		Email:     rec.Email,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}
