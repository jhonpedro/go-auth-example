package userErrors

import "github.com/jhonpedro/go-auth-example/src/shared"

func InvalidEmail() *shared.InternalError {
	return &shared.InternalError{
		StatusCode: 400,
		Message:    "Invalid email",
	}
}

func InvalidUserNameLength() *shared.InternalError {
	return &shared.InternalError{
		StatusCode: 400,
		Message:    "Name must have more than 3 characters",
	}
}

func InvalidUserPasswordLenght() *shared.InternalError {
	return &shared.InternalError{
		StatusCode: 400,
		Message:    "Password must have 8 or more characters",
	}
}

func WrongUserEmailOrPassword() *shared.InternalError {
	return &shared.InternalError{
		StatusCode: 403,
		Message:    "Wrong password or email",
	}
}
