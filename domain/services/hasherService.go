package services

type HasherService interface {
	Create(str string) (string, error)
}
