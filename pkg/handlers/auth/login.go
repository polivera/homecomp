package auth

import (
	"context"
	"homecomp/pkg/templates"
	"net/http"
)

type LoginHandler interface {
	Handle(mux *http.ServeMux)
	showLoginForm(w http.ResponseWriter, _ *http.Request)
}

type loginHaddler struct{}

func NewLoginHandler() LoginHandler {
	return &loginHaddler{}
}

// Login implements LoginHandler.
func (l *loginHaddler) Handle(mux *http.ServeMux) {
	mux.HandleFunc("GET /login", l.showLoginForm)
}

func (l *loginHaddler) showLoginForm(w http.ResponseWriter, _ *http.Request) {
	innerContent := templates.Testing("aijfosaijdfosaijdfoijs")
	component := templates.Layout(innerContent)
	component.Render(context.Background(), w)
}
