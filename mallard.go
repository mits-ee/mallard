package mallard

import (
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	AuthClient *auth.Client
}

func New(config Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token, err := getToken(c, config.AuthClient)
		if err != nil {
			return c.SendStatus(http.StatusUnauthorized)
		}

		if token != nil {
			c.Locals("idToken", token)
		}

		return c.Next()
	}
}

func Perms(c *fiber.Ctx, permissions ...string) int {
	tokenAny := c.Locals("idToken")
	if tokenAny == nil {
		return http.StatusUnauthorized
	}
	token := tokenAny.(*auth.Token)

	if len(permissions) > 0 && !hasPermissions(token, permissions...) {
		return http.StatusForbidden
	}

	return 0
}
