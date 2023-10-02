package database

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var FirestoreClient *firestore.Client

func InitializeFirestore() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("credentials.json")
	client, err := firestore.NewClient(ctx, "gigs-dev-9ab8a", opt)
	if err != nil {
		log.Fatalf("Error initializing Firestore: %v", err)
	}

	FirestoreClient = client
}
