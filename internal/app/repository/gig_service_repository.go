package repository

import (
	"context"
	"github.com/cetalogy/internal/app/model"
	"github.com/cetalogy/internal/database"
	"google.golang.org/api/iterator"
	"log"
)

type GigServiceRepository struct{}

func NewGigServiceRepository() *GigServiceRepository {
	return &GigServiceRepository{}
}

func (ur *GigServiceRepository) GetGigServiceByID(ctx context.Context, ID string) (*model.GigService, error) {
	var result model.GigService

	ref := database.FirestoreClient.Collection("gig_services").Doc(ID)

	snapshot, err := ref.Get(ctx)
	if err != nil {
		log.Printf("Error retrieving gig service: %v", err)
		return nil, err
	}

	if err := snapshot.DataTo(&result); err != nil {
		log.Printf("Error converting gig service data: %v", err)
		return nil, err
	}

	result.ID = ID
	return &result, nil
}

func (ur *GigServiceRepository) GetGigServices(ctx context.Context) (*[]model.GigService, error) {
	var results []model.GigService

	iter := database.FirestoreClient.Collection("gig_services").Documents(ctx)
	for {
		snapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to fetch gig service: %v", err)
			return nil, err
		}

		var result model.GigService
		if err := snapshot.DataTo(&result); err != nil {
			log.Printf("Error converting gig service data: %v", err)
			return nil, err
		}
		results = append(results, result)
	}
	return &results, nil
}

func (ur *GigServiceRepository) CreateGigService(ctx context.Context, req *model.GigService) error {
	ref := database.FirestoreClient.Collection("gig_services").Doc(req.ID)
	_, err := ref.Set(ctx, req)
	if err != nil {
		log.Printf("Error creating gig service: %v", err)
		return err
	}

	return nil
}

func (ur *GigServiceRepository) DeleteGigServiceByID(ctx context.Context, ID string) error {
	ref := database.FirestoreClient.Collection("gig_services").Doc(ID)
	_, err := ref.Delete(ctx)
	if err != nil {
		log.Printf("Error deleting gig service: %v", err)
		return err
	}

	return nil
}
