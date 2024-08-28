package handlers

import (
	"fmt"
	"net/http"
	"time"

	"homecomp/internal/configs"
	"homecomp/internal/database"
	loginform "homecomp/internal/forms/login"
	"homecomp/internal/repositories"
	logintemplate "homecomp/pkg/web/templates/login"
)

const (
	cookieName string = "hcmpauth"
)

type LoginHandler interface {
	Handle(mux *http.ServeMux)
	showLoginForm(w http.ResponseWriter, _ *http.Request)
}

type loginHaddler struct {
	conf     configs.Config
	memDB    database.InMemoryDB
	userRepo repositories.UserRepo
}

func NewLoginHandler(cnf configs.Config, memDB database.InMemoryDB, userRepo repositories.UserRepo) LoginHandler {
	return &loginHaddler{
		conf:     cnf,
		memDB:    memDB,
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

	dbUser := l.userRepo.GetUserByEmail(r.Context(), form.Email)
	if dbUser == nil {
		form.Errors[loginform.FieldEmail] = "User not found "
		component := logintemplate.LoginForm(form)
		component.Render(r.Context(), w)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: "something random", //TODO: Write a randomizer
		Path:  "/",
		// Domain:   "",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	fmt.Fprintln(w, dbUser)
}
