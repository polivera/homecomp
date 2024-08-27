package handlers

import (
	"context"
	"net/http"

	"homecomp/internal/configs"
	"homecomp/internal/database"
	loginform "homecomp/internal/forms/login"
	"homecomp/internal/repositories"
	logintemplate "homecomp/pkg/web/templates/login"
)

type LoginHandler interface {
	Handle(mux *http.ServeMux)
	showLoginForm(w http.ResponseWriter, _ *http.Request)
}

type loginHaddler struct {
	conf      configs.Config
	db        database.DBCon
	ctx       context.Context
	cancelCtx context.CancelFunc
}

func NewLoginHandler(ctx context.Context, cnf configs.Config, db database.DBCon) LoginHandler {
	ctx, cancel := context.WithCancel(ctx)
	return &loginHaddler{
		conf:      cnf,
		db:        db,
		ctx:       ctx,
		cancelCtx: cancel,
	}
}

func (l *loginHaddler) Handle(mux *http.ServeMux) {
	mux.HandleFunc("GET /login", l.showLoginForm)
	mux.HandleFunc("POST /login", l.loginSubmit)
}

func (l *loginHaddler) showLoginForm(w http.ResponseWriter, r *http.Request) {
	component := logintemplate.LoginPage(l.conf.Page)
	component.Render(l.ctx, w)
}

func (l *loginHaddler) loginSubmit(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewUserRepo(r.Context(), l.db)
	form := loginform.LoginForm{}
	form.Email = r.FormValue(loginform.FieldEmail)
	form.Passwd = r.FormValue(loginform.FieldPassword)
	form.Validate(repo)

	if !form.IsValid() {
		component := logintemplate.LoginForm(form)
		component.Render(l.ctx, w)
		return
	}

	w.Write([]byte("subtitution, loosing all my illusion"))
}
