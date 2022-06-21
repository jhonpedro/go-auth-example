package entities_test

import (
	"errors"
	"testing"

	"github.com/jhonpedro/go-auth-example/domain/entities"
	"github.com/jhonpedro/go-auth-example/test/mocks/domainServicesTest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	fakeHasherServiceMock := new(domainServicesTest.FakeHasherService)

	fakeHasherServiceMock.On("Create", mock.Anything).Return("str", nil)

	userEntity := entities.NewUser(fakeHasherServiceMock)

	_, err := userEntity.CreateUser("1", "Jhon", "pedrobarros2010@gmail.com", "123xyz")

	assert.Nil(t, err)
}

func TestNotCreateUserIfHashFails(t *testing.T) {

	fakeHasherServiceMock := new(domainServicesTest.FakeHasherService)

	userEntity := entities.NewUser(fakeHasherServiceMock)

	fakeHasherServiceMock.On("Create", mock.Anything).Return("", errors.New("random error"))

	_, err := userEntity.CreateUser("1", "Jhon", "pedrobarros2010@gmail.com", "123xyz")

	assert.NotNil(t, err)
}
