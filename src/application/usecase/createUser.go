package usecase

import (
	"github.com/jhonpedro/go-auth-example/src/application/dto/userDto"
	"github.com/jhonpedro/go-auth-example/src/domain/entities"
	"github.com/jhonpedro/go-auth-example/src/domain/repositories"
	"github.com/jhonpedro/go-auth-example/src/domain/services"
	"github.com/jhonpedro/go-auth-example/src/shared"
)

type CreateUserUseCase struct {
	uniqueIdentifier services.UniqueIdentifierService
	userRepository   repositories.UserRepository
}

func NewCreateUserUseCase(uniqueIdentifierService services.UniqueIdentifierService, userRepository repositories.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		uniqueIdentifier: uniqueIdentifierService,
		userRepository:   userRepository,
	}
}

func (u *CreateUserUseCase) Execute(input userDto.CreateUserDto) (*entities.User, *shared.InternalError) {
	return nil, nil
}
