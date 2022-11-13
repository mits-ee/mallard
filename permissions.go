package mallard

import "firebase.google.com/go/v4/auth"

const PermissionsKey = "permissions"

func hasPermissions(token *auth.Token, permissions ...string) bool {
	claims, ok := token.Claims[PermissionsKey]
	if !ok {
		return false
	}

	permissionsMap, ok := claims.(map[string]bool)
	if !ok {
		return false
	}

	for _, permission := range permissions {
		if !permissionsMap[permission] {
			return false
		}
	}

	return true
}
