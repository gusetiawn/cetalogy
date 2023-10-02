package repository

import (
	"context"
	"github.com/cetalogy/internal/app/model"
	"github.com/cetalogy/internal/database"
	"google.golang.org/api/iterator"
	"log"
)

type PackageRepository struct{}

func NewPackageRepository() *PackageRepository {
	return &PackageRepository{}
}

func (ur *PackageRepository) GetPackageByID(ctx context.Context, ID string) (*model.Package, error) {
	var result model.Package

	ref := database.FirestoreClient.Collection("packages").Doc(ID)

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

func (ur *PackageRepository) GetPackages(ctx context.Context) (*[]model.Package, error) {
	var results []model.Package

	iter := database.FirestoreClient.Collection("packages").Documents(ctx)
	for {
		snapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to fetch packages: %v", err)
			return nil, err
		}

		var result model.Package
		if err := snapshot.DataTo(&result); err != nil {
			log.Printf("Error converting packages data: %v", err)
			return nil, err
		}
		results = append(results, result)
	}
	return &results, nil
}

func (ur *PackageRepository) CreatePackage(ctx context.Context, req *model.Package) error {
	ref := database.FirestoreClient.Collection("packages").Doc(req.ID)
	_, err := ref.Set(ctx, req)
	if err != nil {
		log.Printf("Error creating packages: %v", err)
		return err
	}

	return nil
}

func (ur *PackageRepository) DeletePackageByID(ctx context.Context, ID string) error {
	ref := database.FirestoreClient.Collection("packages").Doc(ID)
	_, err := ref.Delete(ctx)
	if err != nil {
		log.Printf("Error deleting package: %v", err)
		return err
	}

	return nil
}
