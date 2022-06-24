package usecase

import (
	"github.com/jhonpedro/go-auth-example/src/application/dto/userDto"
	"github.com/jhonpedro/go-auth-example/src/domain/entities"
	"github.com/jhonpedro/go-auth-example/src/domain/repositories"
	"github.com/jhonpedro/go-auth-example/src/domain/services"
	"github.com/jhonpedro/go-auth-example/src/shared"
)

type CreateUserUseCase struct {
	userRepository   repositories.UserRepository
	uniqueIdentifier services.UniqueIdentifierService
	emailValidator   services.EmailValidationService
	hasher           services.HasherService
}

func NewCreateUserUseCase(
	uniqueIdentifierService services.UniqueIdentifierService,
	userRepository repositories.UserRepository,
	emailValidatorService services.EmailValidationService,
	hasherService services.HasherService,
) *CreateUserUseCase {
	return &CreateUserUseCase{
		uniqueIdentifier: uniqueIdentifierService,
		userRepository:   userRepository,
		emailValidator:   emailValidatorService,
		hasher:           hasherService,
	}
}

func (u *CreateUserUseCase) Execute(input userDto.CreateUserDto) (*entities.User, *shared.InternalError) {
	return nil, nil
}
