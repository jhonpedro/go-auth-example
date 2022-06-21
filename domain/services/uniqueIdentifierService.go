package services

type UniqueIdentifierService interface {
	Generate() (string, error)
}
