package repository

import (
	"context"
	"github.com/cetalogy/internal/app/model"
	"github.com/cetalogy/internal/database"
	"google.golang.org/api/iterator"
	"log"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) GetUserByID(ctx context.Context, userID string) (*model.User, error) {
	var result model.User

	ref := database.FirestoreClient.Collection("users").Doc(userID)

	snapshot, err := ref.Get(ctx)
	if err != nil {
		log.Printf("Error retrieving user: %v", err)
		return nil, err
	}

	if err := snapshot.DataTo(&result); err != nil {
		log.Printf("Error converting user data: %v", err)
		return nil, err
	}

	result.ID = userID
	return &result, nil
}

func (ur *UserRepository) GetUsers(ctx context.Context) (*[]model.User, error) {
	var results []model.User

	iter := database.FirestoreClient.Collection("users").Documents(ctx)
	for {
		snapshot, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Printf("Failed to fetch users: %v", err)
			return nil, err
		}

		var result model.User
		if err := snapshot.DataTo(&result); err != nil {
			log.Printf("Error converting user data: %v", err)
			return nil, err
		}
		results = append(results, result)
	}
	return &results, nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	ref := database.FirestoreClient.Collection("users").Doc(user.ID)
	_, err := ref.Set(ctx, user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		return err
	}

	return nil
}

func (ur *UserRepository) DeleteUserByID(ctx context.Context, userID string) error {
	ref := database.FirestoreClient.Collection("users").Doc(userID)
	_, err := ref.Delete(ctx)
	if err != nil {
		log.Printf("Error deleting user: %v", err)
		return err
	}

	return nil
}
