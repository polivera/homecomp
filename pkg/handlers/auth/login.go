package auth

import (
	"context"
	"fmt"
	"net/http"

	"homecomp/internal/configs"
	logintemplate "homecomp/pkg/templates/login"
)

type LoginHandler interface {
	Handle(mux *http.ServeMux)
	showLoginForm(w http.ResponseWriter, _ *http.Request)
}

type loginHaddler struct {
	conf configs.Config
	ctx  context.Context
}

func NewLoginHandler(cnf configs.Config, ctx context.Context) LoginHandler {
	return &loginHaddler{
		conf: cnf,
		ctx:  ctx,
	}
}

// Login implements LoginHandler.
func (l *loginHaddler) Handle(mux *http.ServeMux) {
	mux.HandleFunc("GET /login", l.showLoginForm)
	mux.HandleFunc("POST /login", l.loginSubmit)
	mux.HandleFunc("GET /garompeta", l.garompeta)
}

func (l *loginHaddler) garompeta(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("asofjsaoijfdsaf"))

}

func (l *loginHaddler) showLoginForm(w http.ResponseWriter, _ *http.Request) {
	component := logintemplate.LoginPage(l.conf.Page)
	component.Render(l.ctx, w)
}

func (l *loginHaddler) loginSubmit(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue(logintemplate.EmailField)
	passwd := r.FormValue(logintemplate.PasswordField)
	fmt.Println(email)
	component := logintemplate.LoginForm(
		logintemplate.LoginFormFields{
			Email: passwd,
		},
	)
	component.Render(l.ctx, w)

	// w.Header().Add("HX-Redirect", "/garompeta")
}
