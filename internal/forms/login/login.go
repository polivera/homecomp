package forms

import (
	"homecomp/internal/repositories"
	"homecomp/internal/validators"
)

const (
	FieldEmail    string = "email"
	FieldPassword string = "password"
)

type LoginForm struct {
	Email   string
	Passwd  string
	Errors  map[string]string
	isValid bool
}

func (lf *LoginForm) Validate(repo repositories.UserRepo) {
	lf.isValid = true
	lf.Errors = make(map[string]string, 2)

	if !validators.IsEmailStringValid(lf.Email) {
		lf.Errors[FieldEmail] = "Email address is invalid"
		lf.isValid = false
	} else if !validators.IsEmailNew(lf.Email, repo) {
		lf.Errors[FieldEmail] = "This email is already taken"
		lf.isValid = false
	}

	if validators.IsPasswordLenValid(lf.Passwd) {
		lf.Errors[FieldPassword] = "Password should be at least 8 chars long"
		lf.isValid = false
	}
}

func (lf *LoginForm) IsValid() bool {
	return lf.isValid
}
