package entities

import (
	"time"

	"github.com/jhonpedro/go-auth-example/domain/services"
)

type User struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`

	hasher services.HasherService
}

func NewUser(hasherService services.HasherService) User {
	return User{
		hasher: hasherService,
	}
}

func (u *User) CreateUser(id string, name string, email string, password string) (*User, error) {
	hashedPassword, err := u.hasher.Create(password)

	if err != nil {
		return nil, err
	}

	return &User{
		ID:         id,
		Name:       name,
		Email:      email,
		Password:   hashedPassword,
		Created_at: time.Now(),
	}, nil
}
