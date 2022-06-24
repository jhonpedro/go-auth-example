package services

type EmailValidationService interface {
	IsValid(email string) bool
}
