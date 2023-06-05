# Mallard

MITS' authentication middleware built on top of Firebase Authentication

## Usage

### Installation

`go get github.com/mits-ee/mallard`

### Perms

Function `Perms` can be used to protect routes by enforcing users to have custom claims
passed into the function arguments.

#### Usage

```go
package main

import (
  "firebase.google.com/go/v4/auth"
  mallard "github.com/mits-ee/mallard"
  "github.com/gofiber/fiber/v2"
)

func privateRoute(c *fiber.Ctx) error {
  // This route is protected by Mallard and can only be accessed if the id token
  // contains 'private.has' permission in the custom claims.
  // Read about custom claims here: https://firebase.google.com/docs/auth/admin/custom-claims
  // Mallard will also let requests bypass this when these conditions are met
  // 1. Mallard is created with `mallard.WithBypassHeader` configuration option
  // 2. the request contains the specified header with the given api key
  if err := mallard.Perms(c, "private.has"); err != nil {
			return c.Status(http.StatusUnauthorized).
				SendString(err.Error())
  }

  return c.SendString("Hello, private route!")
}

func main() {
  // Aquire auth client from Firebase (check test/main.go)
  var authClient *auth.Client

  const BypassHeader = "X-Bypass"
  const Apikey = "12345"

  app := fiber.New()
	app.Use(mallard.New(mallard.WithAuthClient(authClient), mallard.WithBypassHeader(BypassHeader, Apikey)))
  app.Get("/private", privateRoute)
  app.Listen(":8080")
}
```

### GetFirestore

Function `GetFirestore` can be used to initialize all needed Firebase services.

#### Usage

```go
package main

import (
  "context"
  "github.com/mits-ee/mallard"
  "log"
)

func main() {
  ctx := context.Background()
  bundle, err := mallard.GetFirebase(ctx, mallard.Authentication)
  if err != nil {
    log.Fatalf("Could not initialize Firebase services: %v\n", err)
  }
  // Now Firebase Authentication can be used
  auth := bundle.Authentication
  // ...
}
```
