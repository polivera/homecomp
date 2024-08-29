package handlers

import (
	"net/http"

	"homecomp/internal/configs"
	"homecomp/internal/database"
	"homecomp/internal/utils"
	dashboardtemplate "homecomp/pkg/web/templates/dashboard"
)

type DashboardHandler interface {
	Handle(mux *http.ServeMux)
	showDashboard(w http.ResponseWriter, _ *http.Request)
}

type dashboardHandler struct {
	conf  configs.Config
	memDB database.InMemoryDB
}

func NewDashboardHandler(cnf configs.Config, memDB database.InMemoryDB) DashboardHandler {
	return &dashboardHandler{
		conf:  cnf,
		memDB: memDB,
	}
}

func (dh *dashboardHandler) Handle(mux *http.ServeMux) {
	mux.HandleFunc("GET /dashboard", utils.UserLoggedMiddleware(dh.showDashboard, dh.memDB))
}

func (dh *dashboardHandler) showDashboard(w http.ResponseWriter, r *http.Request) {
	dashboardtemplate.DashboardPage(dh.conf.Page).Render(r.Context(), w)
}
