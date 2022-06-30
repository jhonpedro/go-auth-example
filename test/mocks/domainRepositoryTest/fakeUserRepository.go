package domainrepositoryTest

import (
	"github.com/jhonpedro/go-auth-example/src/domain/entities"
	"github.com/jhonpedro/go-auth-example/src/shared"
)

type FakeUserRepository struct{}

func (f *FakeUserRepository) Save(user entities.User) (*entities.User, *shared.InternalError) {
	return nil, nil
}

func (f *FakeUserRepository) FindOneByEmail(email string) (*entities.User, *shared.InternalError) {
	return nil, nil
}
