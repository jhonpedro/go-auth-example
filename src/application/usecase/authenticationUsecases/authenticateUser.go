package authenticationUsecases

import (
	"github.com/jhonpedro/go-auth-example/src/application/dto/authenticationDto"
	"github.com/jhonpedro/go-auth-example/src/domain/errors/userErrors"
	"github.com/jhonpedro/go-auth-example/src/domain/repositories"
	"github.com/jhonpedro/go-auth-example/src/domain/services"
	"github.com/jhonpedro/go-auth-example/src/shared"
)

type AuthenticateUserUseCase struct {
	hasher         services.HasherService
	authentication services.AuthenticationService
	userRepository repositories.UserRepository
}

func NewAuthenticateService(
	hasherService services.HasherService,
	authenticationService services.AuthenticationService,
) AuthenticateUserUseCase {
	return AuthenticateUserUseCase{
		hasher:         hasherService,
		authentication: authenticationService,
	}
}

func (a *AuthenticateUserUseCase) Execute(
	input authenticationDto.InputAuthenticateDto,
) (*authenticationDto.OutputAuthenticateDto, *shared.InternalError) {
	usr, err := a.userRepository.FindOneByEmail(input.Email)
	if err != nil {
		return nil, userErrors.WrongUserEmailOrPassword()
	}

	isPasswordRight := a.hasher.Compare(input.Password, usr.Password)

	if !isPasswordRight {
		return nil, userErrors.WrongUserEmailOrPassword()
	}

	token, errToken := a.authentication.CreateJWT(services.InputAuthenticationCreateJWT{
		Id:    usr.ID,
		Email: usr.Email,
	})

	if errToken != nil {
		return nil, errToken
	}

	return &authenticationDto.OutputAuthenticateDto{
		Token: token,
	}, nil
}
