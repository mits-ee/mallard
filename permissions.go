package mallard

import "firebase.google.com/go/v4/auth"

const PermissionsKey = "permissions"

func hasPermissions(token *auth.Token, permissions ...string) bool {
	claims, ok := token.Claims[PermissionsKey]
	if !ok {
		return false
	}

	permissionsMap, ok := claims.(map[string]any)
	if !ok {
		return false
	}

	for _, permission := range permissions {
		if value, ok := permissionsMap[permission]; !ok {
			// Permission is not in the permissions map on the token
			return false
		} else {
			boolValue, ok := value.(bool)
			if !ok || !boolValue {
				// Value is not a boolean or is false
				return false
			}
		}
	}

	return true
}
