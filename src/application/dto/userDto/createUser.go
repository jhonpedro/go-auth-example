package userDto

import "github.com/jhonpedro/go-auth-example/src/domain/entities"

type InputCreateUserDto struct {
	Name       string
	Email      string
	Password   string
	Created_at string
}

type OutputCreateUserDto struct {
	User entities.User
}
