package services

import "github.com/jhonpedro/go-auth-example/src/shared"

type HasherService interface {
	Create(str string) (string, *shared.InternalError)
	Compare(str string, hashedStr string) bool
}
