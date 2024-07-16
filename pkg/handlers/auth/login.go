package auth

import (
	"context"
	"net/http"

	"gitlab.com/xapitan/homecomp/pkg/templates"
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
	innerContent := templates.Testing("mr. cat")
	component := templates.Layout(innerContent, "This is some good title", "en")
	component.Render(context.Background(), w)
}
