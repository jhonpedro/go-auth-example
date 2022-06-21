package domainServicesTest

import "github.com/stretchr/testify/mock"

type FakeHasherService struct {
	mock.Mock
}

func (f *FakeHasherService) Create(str string) (string, error) {
	args := f.Called(str)
	return args.String(0), args.Error(1)
}
