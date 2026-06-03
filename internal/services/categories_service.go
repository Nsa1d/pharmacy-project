package services

import (
	"pharmacy-project/internal/models"
	"pharmacy-project/internal/repository"
)

type CategoryService interface {
	GetAll() ([]models.Category, error)

	GetSubCategoryByCat(categoryID uint) ([]models.SubCategory, error)

	GetMedByCategory(categoryID uint) ([]models.Medicine, error)

	CreateCategory(req models.CategoryUpsert) (*models.Category, error)

	CreateSubCategory(categoryID uint, req models.SubCategoryUpsert) (*models.SubCategory, error)
}

type categoryService struct {
	category repository.CategoryRepository
}

func NewCategoryService(service repository.CategoryRepository) CategoryService {
	return &categoryService{category: service}
}

func (s *categoryService) GetAll() ([]models.Category, error) {
	cat, err := s.category.GetAllCat()
	if err != nil {
		return nil, err
	}
	return cat, nil
}

func (s *categoryService) GetSubCategoryByCat(categoryID uint) ([]models.SubCategory, error) {
	cat, err := s.category.GetSubCategoryByCat(categoryID)
	if err != nil {
		return nil, err
	}
	return cat, nil
}

func (s *categoryService) GetMedByCategory(categoryID uint) ([]models.Medicine, error) {
	med, err := s.category.GetMedByCategory(categoryID)
	if err != nil {
		return nil, err
	}
	return med, nil
}

func (s *categoryService) CreateCategory(req models.CategoryUpsert) (*models.Category, error) {
	category := &models.Category{Name: req.Name}
	if err := s.category.CreateCategory(category); err != nil {
		return nil, err
	}
	return category, nil
}

func (s *categoryService) CreateSubCategory(categoryID uint, req models.SubCategoryUpsert) (*models.SubCategory, error) {
	if _, err := s.category.GetSubCategoryByCat(categoryID); err != nil {
		return nil, err
	}
	subCategory := &models.SubCategory{
		Name:       req.Name,
		CategoryID: categoryID,
	}
	if err := s.category.CreateSubCategory(subCategory); err != nil {
		return nil, err
	}
	return subCategory, nil
}
