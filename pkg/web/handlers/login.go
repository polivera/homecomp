package handlers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"homecomp/internal/configs"
	"homecomp/internal/database"
	loginform "homecomp/internal/forms/login"
	"homecomp/internal/repositories"
	"homecomp/internal/utils"
	logintemplate "homecomp/pkg/web/templates/login"
)

type LoginHandler interface {
	Handle(mux *http.ServeMux)
	showLoginForm(w http.ResponseWriter, _ *http.Request)
}

type loginHandler struct {
	conf     configs.Config
	memDB    database.InMemoryDB
	userRepo repositories.UserRepo
}

func NewLoginHandler(cnf configs.Config, memDB database.InMemoryDB, userRepo repositories.UserRepo) LoginHandler {
	return &loginHandler{
		conf:     cnf,
		memDB:    memDB,
		userRepo: userRepo,
	}
}

func (l *loginHandler) Handle(mux *http.ServeMux) {
	mux.HandleFunc("GET /login", l.showLoginForm)
	mux.HandleFunc("POST /login", l.loginSubmit)
}

func (l *loginHandler) showLoginForm(w http.ResponseWriter, r *http.Request) {
	if _, err := utils.LoggedInUser(r, l.memDB); err == nil {
		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
		return
	}
	component := logintemplate.LoginPage(l.conf.Page)
	component.Render(r.Context(), w)
}

func (l *loginHandler) loginSubmit(w http.ResponseWriter, r *http.Request) {
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
		form.Errors[loginform.FieldEmail] = "Invalid user or password"
		component := logintemplate.LoginForm(form)
		component.Render(r.Context(), w)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(form.Passwd))
	if err != nil {
		form.Errors[loginform.FieldEmail] = "Invalid user or password"
		component := logintemplate.LoginForm(form)
		component.Render(r.Context(), w)
		return
	}

	utils.LoginUser(*dbUser, l.memDB, w, r)

	w.Header().Add("HX-Redirect", "/dashboard")
}
