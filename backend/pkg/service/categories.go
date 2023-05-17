package service

import (
	budget "budget-tracker"
	"budget-tracker/pkg/repository"
)

type CategoriesService struct {
	repo repository.Categories
}

func (cs *CategoriesService) GetAllCategories() (categories []budget.Categories, err error) {
	return cs.repo.GetAllCategories()
}
