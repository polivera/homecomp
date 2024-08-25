package auth

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"homecomp/internal/configs"
	"homecomp/internal/database"
	"homecomp/internal/repositories"
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

	// TODO: validate form data here
	// TODO: Work on form error

	hashPass, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		w.Write([]byte("some error"))
		return
	}
	repo := repositories.NewUserRepo(r.Context(), l.db)

	//TODO: Finish login with a select
	err = repo.CreateUser(repositories.UserRow{Email: email, Password: string(hashPass)})
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	// TODO: How to eat cookies

	w.Write([]byte("done"))

	// component := logintemplate.LoginForm(
	// 	logintemplate.LoginFormFields{
	// 		Email: passwd,
	// 	},
	// )
	// component.Render(l.ctx, w)

	// w.Header().Add("HX-Redirect", "/garompeta")
}
