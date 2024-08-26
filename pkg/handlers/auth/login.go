package auth

import (
	"fmt"
	"net/http"

	"homecomp/internal/configs"
	"homecomp/internal/database"
	logintemplate "homecomp/pkg/templates/login"
)

type LoginHandler interface {
	Handle(mux *http.ServeMux)
	showLoginForm(w http.ResponseWriter, _ *http.Request)
}

type loginHaddler struct {
	conf configs.Config
	db   database.DBCon
}

func NewLoginHandler(cnf configs.Config, db database.DBCon) LoginHandler {
	return &loginHaddler{
		conf: cnf,
		db:   db,
	}
}

// Login implements LoginHandler.
func (l *loginHaddler) Handle(mux *http.ServeMux) {
	mux.HandleFunc("GET /login", l.showLoginForm)
	mux.HandleFunc("POST /login", l.loginSubmit)
}

func (l *loginHaddler) showLoginForm(w http.ResponseWriter, r *http.Request) {
	component := logintemplate.LoginPage(l.conf.Page)
	component.Render(r.Context(), w)
}

func (l *loginHaddler) loginSubmit(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue(logintemplate.EmailField)
	passwd := r.FormValue(logintemplate.PasswordField)
	fmt.Println(passwd)

	// TODO: validate form data here

	component := logintemplate.LoginForm(
		logintemplate.LoginFormFields{
			Email:         email,
			EmailError:    "some email error",
			PasswordError: "some password error",
		},
	)
	component.Render(r.Context(), w)
	// TODO: Work on form error

	// hashPass, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	// if err != nil {
	// 	w.Write([]byte("some error"))
	// 	return
	// }
	// TODO: How to eat cookies

	// w.Write([]byte("done"))

	// component := logintemplate.LoginForm(
	// 	logintemplate.LoginFormFields{
	// 		Email: passwd,
	// 	},
	// )
	// component.Render(l.ctx, w)

	// w.Header().Add("HX-Redirect", "/garompeta")
}
