package mallard

import (
	"cloud.google.com/go/firestore"
	"context"
	"errors"
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
	"log"
	"os"
	"strings"
)

type FirebaseServiceType int

const (
	Authentication FirebaseServiceType = iota
	Firestore
)

type ServiceBundle struct {
	App            *firebase.App
	Authentication *auth.Client
	Firestore      *firestore.Client
}

var ErrorAlreadyInitialized = errors.New("service is already initialized")
var ErrorUnsupportedService = errors.New("unsupported service")

func GetFirebase(ctx context.Context, services ...FirebaseServiceType) (*ServiceBundle, error) {
	bundle := &ServiceBundle{}
	var err error

	bundle.App, err = getApp()
	if err != nil {
		return nil, err
	}

	for _, service := range services {
		switch service {
		case Authentication:
			if bundle.Authentication != nil {
				return nil, ErrorAlreadyInitialized
			}

			bundle.Authentication, err = bundle.App.Auth(ctx)
			if err != nil {
				return nil, ErrorAlreadyInitialized
			}
		case Firestore:
			if bundle.Firestore != nil {
				return nil, ErrorAlreadyInitialized
			}

			bundle.Firestore, err = bundle.App.Firestore(ctx)
			if err != nil {
				return nil, ErrorAlreadyInitialized
			}
		default:
			return nil, ErrorUnsupportedService
		}
	}

	return bundle, nil
}

func getApp() (*firebase.App, error) {
	ctx := context.Background()
	env := strings.ToUpper(os.Getenv("MITS_ENV"))
	log.Printf("Initializing Firebase in environment %s\n", env)

	if env == "DEV" {
		serviceAccountJsonLocation := os.Getenv("MITS_SERVICE_ACCOUNT")
		if serviceAccountJsonLocation == "" {
			serviceAccountJsonLocation = "./service_account.json"
		}
		log.Printf("Using service account %s\n", serviceAccountJsonLocation)

		opt := option.WithCredentialsFile(serviceAccountJsonLocation)
		return firebase.NewApp(ctx, nil, opt)
	}

	return firebase.NewApp(ctx, nil)
}
