package entities

import (
	"time"

	"github.com/jhonpedro/go-auth-example/src/domain/errors/userErrors"
	"github.com/jhonpedro/go-auth-example/src/domain/services"
	"github.com/jhonpedro/go-auth-example/src/shared"
)

type User struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_at time.Time `json:"created_at"`

	hasher         services.HasherService
	emailValidator services.EmailValidationService
}

func NewUser(hasherService services.HasherService, emailValidatorService services.EmailValidationService) User {
	return User{
		hasher:         hasherService,
		emailValidator: emailValidatorService,
	}
}

func (u *User) CreateUser(id string, name string, email string, password string) (*User, *shared.InternalError) {
	emailIsValid := u.emailValidator.IsValid(email)

	if !emailIsValid {
		return nil, userErrors.InvalidEmail()
	}

	if len(name) < 3 {
		return nil, userErrors.InvalidUserNameLength()
	}

	if len(password) < 8 {
		return nil, userErrors.InvalidUserPasswordLenght()
	}

	hashedPassword, err := u.hasher.Create(password)
	if err != nil {
		return nil, err
	}

	u.ID = id
	u.Name = name
	u.Email = email
	u.Password = hashedPassword
	u.Created_at = time.Now()

	return u, nil
}
