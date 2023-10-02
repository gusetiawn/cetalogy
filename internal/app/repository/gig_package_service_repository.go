package repository

import (
	"context"
	"github.com/cetalogy/internal/app/model"
	"github.com/cetalogy/internal/database"
	"google.golang.org/api/iterator"
	"log"
)

type GigPackageServiceRepository struct{}

func NewGigPackageServiceRepository() *GigPackageServiceRepository {
	return &GigPackageServiceRepository{}
}

func (ur *GigPackageServiceRepository) GetGigPackageServiceByID(ctx context.Context, ID string) (*model.GigPackageService, error) {
	var result model.GigPackageService

	ref := database.FirestoreClient.Collection("gig_package_services").Doc(ID)

	snapshot, err := ref.Get(ctx)
	if err != nil {
		log.Printf("Error retrieving Package: %v", err)
		return nil, err
	}

	if err := snapshot.DataTo(&result); err != nil {
		log.Printf("Error converting Package data: %v", err)
		return nil, err
	}

	result.ID = ID
	return &result, nil
}

func (ur *GigPackageServiceRepository) GetGigPackageServices(ctx context.Context) (*[]model.GigPackageService, error) {
	var results []model.GigPackageService

	iter := database.FirestoreClient.Collection("gig_package_services").Documents(ctx)
	for {
		snapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to fetch packages: %v", err)
			return nil, err
		}

		var result model.GigPackageService
		if err := snapshot.DataTo(&result); err != nil {
			log.Printf("Error converting packages data: %v", err)
			return nil, err
		}
		results = append(results, result)
	}
	return &results, nil
}

func (ur *GigPackageServiceRepository) CreateGigPackageService(ctx context.Context, req *model.GigPackageService) error {
	ref := database.FirestoreClient.Collection("gig_package_services").Doc(req.ID)
	_, err := ref.Set(ctx, req)
	if err != nil {
		log.Printf("Error creating packages: %v", err)
		return err
	}

	return nil
}

func (ur *GigPackageServiceRepository) DeleteGigPackageServiceByID(ctx context.Context, ID string) error {
	ref := database.FirestoreClient.Collection("gig_package_services").Doc(ID)
	_, err := ref.Delete(ctx)
	if err != nil {
		log.Printf("Error deleting package: %v", err)
		return err
	}

	return nil
}
