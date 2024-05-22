package users

import (
	"errors"
	"fmt"
	"mkp-cinema-api/businesses/users"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &userRepository{
		conn: conn,
	}
}

func (ur *userRepository) Register(userDomain *users.Domain) (users.Domain, error) {
	password, _ := bcrypt.GenerateFromPassword([]byte(userDomain.Password), bcrypt.DefaultCost)
	rec := FromDomain(userDomain)
	rec.Password = string(password)

	// Check if email is already used
	var userByEmail User
	if err := ur.conn.First(&userByEmail, "email = ?", userDomain.Email).Error; err == nil {
		return users.Domain{}, errors.New("email already taken, please use another email or process to login")
	}

	// Check if username is already used
	var userByUsername User
	if err := ur.conn.First(&userByUsername, "username = ?", userDomain.Username).Error; err == nil {
		return users.Domain{}, errors.New("username already taken, please choose another username")
	}

	// Create new user
	if err := ur.conn.Create(&rec).Error; err != nil {
		fmt.Println("error creating user:", err)
		return users.Domain{}, errors.New("internal server error")
	}

	return rec.ToDomain(), nil
}

func (ur *userRepository) GetByUsername(userDomain *users.LoginDomain) (users.Domain, error) {
	var user User
	if err := ur.conn.First(&user, "username = ?", userDomain.Username).Error; err != nil {
		return users.Domain{}, errors.New("user not found")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password))
	if err != nil {
		return users.Domain{}, errors.New("password failed")
	}

	return user.ToDomain(), nil
}
