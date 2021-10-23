package service

import (
	shop "github.com/zangar-tm/online_shop"
	"github.com/zangar-tm/online_shop/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) Create(category shop.Category) (int, error) {
	return s.repo.Create(category)
}

func (s *CategoryService) GetAll() ([]shop.Category, error) {
	return s.repo.GetAll()
}

func (s *CategoryService) GetById(categoryId int) (shop.Category, error) {
	return s.repo.GetById(categoryId)
}

func (s *CategoryService) Delete(categoryId int) error {
	return s.repo.Delete(categoryId)
}

// func (s *Service) Update(categoryId int, input shop.UpdateCategoryInput) error {

// }
