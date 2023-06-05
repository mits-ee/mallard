package main

import (
	"context"
	"net/http"

	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/mits-ee/mallard"
)

const BypassHeader = "X-Bypass"
const Apikey = "12345"

func main() {
	var authClient *auth.Client

	ctx := context.Background()
	bundle, err := mallard.GetFirebase(ctx, mallard.Authentication)
	if err != nil {
		panic(err)
	}
	authClient = bundle.Authentication

	app := fiber.New()
	app.Use(mallard.New(mallard.WithAuthClient(authClient), mallard.WithBypassHeader(BypassHeader, Apikey)))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("public")
	})

	app.Get("/private", func(c *fiber.Ctx) error {
		if err := mallard.Perms(c, "private.has"); err != nil {
			return c.Status(http.StatusUnauthorized).
				SendString(err.Error())
		}

		return c.SendString("private")
	})

	app.Listen(":8081")
}
