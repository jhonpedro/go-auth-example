package repositories

import (
	"github.com/jhonpedro/go-auth-example/src/domain/entities"
	"github.com/jhonpedro/go-auth-example/src/shared"
)

type UserRepository interface {
	Save(user entities.User) (*entities.User, *shared.InternalError)
	FindOneByEmail(email string) (*entities.User, *shared.InternalError)
}
