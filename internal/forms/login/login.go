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

	if lf.Passwd == "" {
		lf.Errors[FieldPassword] = "Password is required"
		lf.isValid = false
	} else {
		if !validators.IsEmailStringValid(lf.Email) {
			lf.Errors[FieldEmail] = "Email address is invalid"
			lf.isValid = false
		}
	}

	if lf.Email == "" {
		lf.Errors[FieldEmail] = "Email is required"
		lf.isValid = false
	} else {
		if validators.IsPasswordLenValid(lf.Passwd) {
			lf.Errors[FieldEmail] = "Invalid email or password"
			lf.isValid = false
		} else if !validators.IsPasswordCharsValid(lf.Passwd) {
			lf.Errors[FieldEmail] = "Invalid email or password"
			lf.isValid = false
		}
	}
}

func (lf *LoginForm) IsValid() bool {
	return lf.isValid
}
