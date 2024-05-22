package response

import (
	"mkp-cinema-api/businesses/users"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey"`
	CreatedAt string    `json:"created_at"`
	UpdatedAt string    `json:"updated_at"`
	DeletedAt string    `json:"deleted_at"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:        domain.ID,
		Username:  domain.Username,
		Password:  domain.Password,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt.Format("02-01-2006 15:04:05"),
		UpdatedAt: domain.UpdatedAt.Format("02-01-2006 15:04:05"),
		DeletedAt: domain.DeletedAt.Time.Format("02-01-2006 15:04:05"),
	}
}
