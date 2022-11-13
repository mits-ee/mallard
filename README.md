# Mallard

MITS' authentication middleware built on top of Firebase Authentication

## Usage

### Installation

`go get github.com/mits-ee/mallard`

### Usage example using Gorilla Mux

```go
package main

import (
	"firebase.google.com/go/v4/auth"
	mallard "github.com/mits-ee/mallard"
	"github.com/gorilla/mux"
	"net/http"
)

func privateRoute() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "you have access to the private route")
    }
}

func main() {
    // Set this up somehow
    var authClient *auth.Client

    router := mux.NewRouter()
    router.HandleFunc("/private", mallard.Perms(authClient, privateRoute(), "private.get")).
        Methods("GET")
		
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Error starting server: %v\n", err)
    }
}
```