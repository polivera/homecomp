package utils

import (
	"net/http"
	"time"

	"homecomp/internal/database"
	"homecomp/internal/repositories"
)

const (
	cookieName string = "hcmpauth"
)

func LoginUser(dbUser repositories.UserRow, memDB database.InMemoryDB, w http.ResponseWriter, r *http.Request) {
	sessionID := RandomStrOfLen(32)
	memDB.Set(r.Context(), sessionID, dbUser.ID)

	http.SetCookie(w, &http.Cookie{
		Name:  cookieName,
		Value: sessionID,
		Path:  "/",
		// Domain:   "",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})
}

func LoggedInUser(r *http.Request, memDB database.InMemoryDB) (uint32, error) {
	sessionCookie, err := r.Cookie(cookieName)
	if err != nil {
		return 0, err
	}
	uID, err := memDB.Get(r.Context(), sessionCookie.Value)
	if err != nil {
		return 0, err
	}
	return uID, nil
}

func UserLoggedMiddleware(next http.HandlerFunc, memDB database.InMemoryDB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := LoggedInUser(r, memDB); err != nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
