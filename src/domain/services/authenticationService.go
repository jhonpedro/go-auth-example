package services

import "github.com/jhonpedro/go-auth-example/src/shared"

type InputAuthenticationCreateJWT struct {
	Id    string
	Email string
}

type AuthenticationService interface {
	CreateJWT(input InputAuthenticationCreateJWT) (string, *shared.InternalError)
}
