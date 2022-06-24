package domainServicesTest

import (
	"github.com/jhonpedro/go-auth-example/src/shared"
	"github.com/stretchr/testify/mock"
)

type FakeHasherService struct {
	mock.Mock
}

func (f *FakeHasherService) Create(str string) (string, *shared.InternalError) {
	args := f.Called(str)

	var error *shared.InternalError

	if args.Get(1) == nil {
		error = nil
	} else {
		error = args.Get(1).(*shared.InternalError)
	}

	return args.String(0), error
}
