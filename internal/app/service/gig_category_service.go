package service

import (
	"context"

	"github.com/cetalogy/internal/app/model"
	"github.com/cetalogy/internal/app/repository"
)

type GigCategoryService struct {
	gigCategoryRepository *repository.GigCategoryRepository
}

func NewGigCategoryService(ur *repository.GigCategoryRepository) *GigCategoryService {
	return &GigCategoryService{gigCategoryRepository: ur}
}

func (us *GigCategoryService) CreateGigCategory(ctx context.Context, user *model.GigServiceCategory) error {
	return us.gigCategoryRepository.CreateCategory(ctx, user)
}

func (us *GigCategoryService) GetGigCategoryByID(ctx context.Context, categoryID string) (*model.GigServiceCategory, error) {
	return us.gigCategoryRepository.GetCategoryByID(ctx, categoryID)
}

func (us *GigCategoryService) GetGigCategories(ctx context.Context) (*[]model.GigServiceCategory, error) {
	return us.gigCategoryRepository.GetCategories(ctx)
}

func (us *GigCategoryService) DeleteGigCategoryByID(ctx context.Context, categoryID string) error {
	return us.gigCategoryRepository.DeleteGigCategoryByID(ctx, categoryID)
}
