package entities_test

import (
	"testing"

	"github.com/jhonpedro/go-auth-example/src/domain/entities"
	"github.com/jhonpedro/go-auth-example/src/domain/errors/userErrors"
	"github.com/jhonpedro/go-auth-example/src/shared"
	"github.com/jhonpedro/go-auth-example/test/mocks/domainServicesTest"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var fakeHasherServiceMock = new(domainServicesTest.FakeHasherService)
var fakeEmailValidator = new(domainServicesTest.EmailValidationService)

func TestCreateUser(t *testing.T) {
	fakeHasherServiceMock.On("Create", mock.Anything).Return("str", nil)
	fakeEmailValidator.On("IsValid", mock.Anything).Return(true)

	userEntity := entities.NewUser(fakeHasherServiceMock, fakeEmailValidator)

	_, err := userEntity.CreateUser("1", "Jhon", "pedrobarros2010@gmail.com", "123xyz456")

	assert.Nil(t, err)
}

func TestNotCreateUserIfEmailIsInvalid(t *testing.T) {
	fakeHasherServiceMock.On("Create", mock.Anything).Return("str", nil)
	fakeEmailValidator.On("IsValid", mock.Anything).Return(false)

	userEntity := entities.NewUser(fakeHasherServiceMock, fakeEmailValidator)

	_, err := userEntity.CreateUser("1", "Jhon", "pedrobarros2010@gmail.com", "123xyz456")

	assert.Nil(t, err)
}

func TestNotCreateUserIfNameIsInvalid(t *testing.T) {
	fakeHasherServiceMock.On("Create", mock.Anything).Return("str", nil)
	fakeEmailValidator.On("IsValid", mock.Anything).Return(true)

	userEntity := entities.NewUser(fakeHasherServiceMock, fakeEmailValidator)

	_, err := userEntity.CreateUser("1", "Yu", "pedrobarros2010@gmail.com", "123xyz456")

	assert.Equal(t, *userErrors.InvalidUserNameLength(), *err)
}

func TestNotCreateUserIfPaswordHasInvalidLength(t *testing.T) {
	fakeHasherServiceMock.On("Create", mock.Anything).Return("str", nil)
	fakeEmailValidator.On("IsValid", mock.Anything).Return(true)

	userEntity := entities.NewUser(fakeHasherServiceMock, fakeEmailValidator)

	_, err := userEntity.CreateUser("1", "Jhon", "pedrobarros2010@gmail.com", "123")

	assert.Equal(t, *userErrors.InvalidUserPasswordLenght(), *err)
}

func TestNotCreateUserIfHashFails(t *testing.T) {
	fakeHasherServiceMock := new(domainServicesTest.FakeHasherService)
	fakeEmailValidator.On("IsValid", mock.Anything).Return(true)

	userEntity := entities.NewUser(fakeHasherServiceMock, fakeEmailValidator)

	fakeHasherServiceMock.On("Create", mock.Anything).Return("", new(shared.InternalError))

	_, err := userEntity.CreateUser("1", "Jhon", "pedrobarros2010@gmail.com", "123xyz456")

	assert.NotNil(t, err)
}
