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
  if status := mallard.Perms(c, "private.has"); status != 0 {
    return c.SendStatus(status)
  }

  return c.SendString("Hello, private route!")
}

func main() {
  // Aquire auth client from Firebase
  var authClient *auth.Client

  app := fiber.New()
  app.Use(mallard.New(mallard.Config{
    AuthClient: authClient,
  }))
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
