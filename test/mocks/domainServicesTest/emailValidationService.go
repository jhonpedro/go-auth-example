package domainServicesTest

import "github.com/stretchr/testify/mock"

type EmailValidationService struct {
	mock.Mock
}

func (e *EmailValidationService) IsValid(email string) bool {
	args := e.Called(email)

	return args.Bool(0)
}
