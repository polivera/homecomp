package handlers

import (
	"net/http"

	"homecomp/internal/configs"
	loginform "homecomp/internal/forms/login"
	"homecomp/internal/repositories"
	logintemplate "homecomp/pkg/web/templates/login"
)

type LoginHandler interface {
	Handle(mux *http.ServeMux)
	showLoginForm(w http.ResponseWriter, _ *http.Request)
}

type loginHaddler struct {
	conf     configs.Config
	userRepo repositories.UserRepo
}

func NewLoginHandler(cnf configs.Config, userRepo repositories.UserRepo) LoginHandler {
	return &loginHaddler{
		conf:     cnf,
		userRepo: userRepo,
	}
}

func (l *loginHaddler) Handle(mux *http.ServeMux) {
	mux.HandleFunc("GET /login", l.showLoginForm)
	mux.HandleFunc("POST /login", l.loginSubmit)
}

func (l *loginHaddler) showLoginForm(w http.ResponseWriter, r *http.Request) {
	component := logintemplate.LoginPage(l.conf.Page)
	component.Render(r.Context(), w)
}

func (l *loginHaddler) loginSubmit(w http.ResponseWriter, r *http.Request) {
	form := loginform.LoginForm{}
	form.Email = r.FormValue(loginform.FieldEmail)
	form.Passwd = r.FormValue(loginform.FieldPassword)
	form.Validate(l.userRepo)

	if !form.IsValid() {
		component := logintemplate.LoginForm(form)
		component.Render(r.Context(), w)
		return
	}

	w.Write([]byte("subtitution, loosing all my illusion"))
}
