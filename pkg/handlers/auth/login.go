package auth

import (
	"context"
	"net/http"

	"homecomp/pkg/configs"
	"homecomp/pkg/templates"
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
}

func (l *loginHaddler) showLoginForm(w http.ResponseWriter, _ *http.Request) {
	component := templates.LoginPage(l.conf)
	component.Render(l.ctx, w)
}
