package repository

import (
	"hobee/model"
)

type CategoryRepository interface {
	GetAllCategory() ([]model.Category, error)
	}

type categoryRepo struct {
}
func NewCategoryRepository() CategoryRepository {
	return &categoryRepo{}
}


	func (s categoryRepo) GetAllCategory() ([]model.Category, error) {
		var categories []model.Category
		err := db.Find(&categories).Error
		if err != nil {
			return categories, err
		}
		return categories, nil
	}