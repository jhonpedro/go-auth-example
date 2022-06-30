package authenticationDto

type InputAuthenticateDto struct {
	Email    string
	Password string
}

type OutputAuthenticateDto struct {
	Token string
}
