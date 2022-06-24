package services

import "github.com/jhonpedro/go-auth-example/src/shared"

type UniqueIdentifierService interface {
	Generate() (string, *shared.InternalError)
}
