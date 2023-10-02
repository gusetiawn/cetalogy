package repository

import (
	"context"
	"github.com/cetalogy/internal/app/model"
	"github.com/cetalogy/internal/database"
	"google.golang.org/api/iterator"
	"log"
)

type GigCategoryRepository struct{}

func NewCategoryRepository() *GigCategoryRepository {
	return &GigCategoryRepository{}
}

func (ur *GigCategoryRepository) GetCategoryByID(ctx context.Context, ID string) (*model.GigServiceCategory, error) {
	var result model.GigServiceCategory

	ref := database.FirestoreClient.Collection("gig_service_categories").Doc(ID)

	snapshot, err := ref.Get(ctx)
	if err != nil {
		log.Printf("Error retrieving gig category: %v", err)
		return nil, err
	}

	if err := snapshot.DataTo(&result); err != nil {
		log.Printf("Error converting gig category data: %v", err)
		return nil, err
	}

	//result.ID = ID
	return &result, nil
}

func (ur *GigCategoryRepository) GetCategories(ctx context.Context) (*[]model.GigServiceCategory, error) {
	var result []model.GigServiceCategory

	iter := database.FirestoreClient.Collection("gig_service_categories").Documents(ctx)
	for {
		snapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to fetch gig service category: %v", err)
			return nil, err
		}

		var category model.GigServiceCategory
		if err := snapshot.DataTo(&category); err != nil {
			log.Printf("Error converting gig service category data: %v", err)
			return nil, err
		}
		result = append(result, category)
	}
	return &result, nil
}

func (ur *GigCategoryRepository) CreateCategory(ctx context.Context, req *model.GigServiceCategory) error {
	ref := database.FirestoreClient.Collection("gig_service_categories").Doc(req.ID)
	_, err := ref.Set(ctx, req)
	if err != nil {
		log.Printf("Error creating gig service category: %v", err)
		return err
	}

	return nil
}

func (ur *GigCategoryRepository) DeleteGigCategoryByID(ctx context.Context, ID string) error {
	ref := database.FirestoreClient.Collection("gig_service_categories").Doc(ID)
	_, err := ref.Delete(ctx)
	if err != nil {
		log.Printf("Error deleting gig category: %v", err)
		return err
	}

	return nil
}
