package service

import (
	"hobee/repository"
	"hobee/model"
)

type CategoryService interface {
	GetAllCategory() ([]model.Category, error)
}
type serviceCategory struct {
}

func NewCategoryService() CategoryService {
	return &serviceCategory{}
}

func (s serviceCategory) GetAllCategory() ([]model.Category, error) {
	var categoryRepo = repository.NewCategoryRepository()
	category, err := categoryRepo.GetAllCategory()
	if err != nil {
		return category, err
	}
	return category, nil
}
