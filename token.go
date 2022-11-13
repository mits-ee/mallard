package mallard

import (
	"fmt"
	"net/http"
	"strings"

	"firebase.google.com/go/v4/auth"
)

type tokenKeyType struct{ string }

var TokenKey = tokenKeyType{"token"}

func getToken(r *http.Request, authClient *auth.Client) (*auth.Token, error) {
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		return nil, fmt.Errorf("token not found")
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

	return authClient.VerifyIDToken(r.Context(), tokenStr)
}
