package mallard

import (
	"context"
	"net/http"

	"firebase.google.com/go/v4/auth"
)

// Perms
// Usage:
//
//	router.HandleFunc("/users", mallard.Perms(authClient, handleUsers(), "users.list")).Methods("GET")
func Perms(authClient *auth.Client, next http.HandlerFunc, permissions ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := getToken(r, authClient)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if len(permissions) > 0 && !hasPermissions(token, permissions...) {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		ctx := context.WithValue(r.Context(), TokenKey, token)
		next(w, r.WithContext(ctx))
	}
}
