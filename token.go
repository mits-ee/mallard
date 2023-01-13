package mallard

import (
	"strings"

	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
)

func getToken(c *fiber.Ctx, authClient *auth.Client) (*auth.Token, error) {
	tokenStr := c.Get(fiber.HeaderAuthorization)
	if tokenStr == "" {
		return nil, nil
	}

	tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")
	return authClient.VerifyIDToken(c.Context(), tokenStr)
}
