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
	userId, errId := u.uniqueIdentifier.Generate()
	if errId != nil {
		return nil, errId
	}

	userEntity := entities.NewUser(u.hasher, u.emailValidator)

	_, errUserEntity := userEntity.CreateUser(userId, input.Name, input.Email, input.Password)

	if errUserEntity != nil {
		return nil, errUserEntity
	}

	savedUser, errSaveUser := u.userRepository.Save(userEntity)

	if errSaveUser != nil {
		return nil, errSaveUser
	}

	return savedUser, nil
}
