package mallard

import (
	"errors"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
)

type Opts struct {
	authClient   *auth.Client
	bypassHeader string
	apiKey       string
}

type OptFunc func(*Opts)

func New(opts ...OptFunc) fiber.Handler {
	baseOpts := new(Opts)
	for _, opt := range opts {
		opt(baseOpts)
	}

	if baseOpts.authClient == nil {
		panic("Firebase authentication client is missing for mallard client")
	}

	return func(c *fiber.Ctx) error {
		c.Locals("mallardOpts", baseOpts)

		token, err := getToken(c, baseOpts.authClient)
		if err != nil {
			return c.SendStatus(http.StatusUnauthorized)
		}

		if token != nil {
			c.Locals("idToken", token)
		}

		return c.Next()
	}
}

func WithAuthClient(authClient *auth.Client) OptFunc {
	return func(o *Opts) {
		o.authClient = authClient
	}
}

func WithBypassHeader(bypassHeader string, apiKey string) OptFunc {
	return func(o *Opts) {
		o.bypassHeader = bypassHeader
		o.apiKey = apiKey
	}
}

func Perms(c *fiber.Ctx, permissions ...string) error {
	mallardOpts := c.Locals("mallardOpts").(*Opts)

	shouldCheckBypass := mallardOpts.bypassHeader != ""

	var bypassErr error
	if shouldCheckBypass {
		bypassErr = checkBypass(c, mallardOpts)
	}

	var tokenErr error
	tokenAny := c.Locals("idToken")
	if tokenAny != nil {
		token := tokenAny.(*auth.Token)

		if len(permissions) > 0 && !hasPermissions(token, permissions...) {
			tokenErr = errors.New("forbidden")
		}
	} else {
		tokenErr = errors.New("unauthorized")
	}

	if tokenErr != nil && !shouldCheckBypass {
		return tokenErr
	}

	return bypassErr
}
