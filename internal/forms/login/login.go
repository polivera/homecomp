package forms

import "homecomp/internal/validators"

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

func (lf *LoginForm) Validate() {
	lf.isValid = true
	lf.Errors = make(map[string]string, 2)
	emailErr := validators.IsEmailValid(lf.Email)
	if emailErr != nil {
		lf.Errors[FieldEmail] = emailErr.Error()
		lf.isValid = false
	}

	passErr := validators.IsValidPassword(lf.Passwd)
	if passErr != nil {
		lf.Errors[FieldPassword] = passErr.Error()
		lf.isValid = false
	}
}

func (lf *LoginForm) IsValid() bool {
	return lf.isValid
}
