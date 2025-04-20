package middleware

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/whynotavailable/svc"
)

type AuthMiddleware struct {
	Inner http.Handler
}

func (a *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" && r.URL.Path == "/_info" {
		// Exlude the docs from requiring a token
		a.Inner.ServeHTTP(w, r)
		return
	}

	val, ok := r.Header["Authorization"]

	if !ok {
		svc.WriteErrorWithCode(w, errors.New("not authorized"), http.StatusUnauthorized)
		return
	}

	if len(val) != 1 {
		slog.Warn("got duplicate authorization headers")
		svc.WriteErrorWithCode(w, errors.New("not authorized"), http.StatusUnauthorized)
		return
	}

	// Simulate token validation
	if val[0] == "good" {
		a.Inner.ServeHTTP(w, r)
	} else {
		svc.WriteErrorWithCode(w, errors.New("not authorized"), http.StatusUnauthorized)
		return
	}
}
