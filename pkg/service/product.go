package service

import (
	shop "github.com/zangar-tm/online_shop"
	"github.com/zangar-tm/online_shop/pkg/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(categoryId int, product shop.Product) (int, error) {
	return s.repo.Create(categoryId, product)
}

func (s *ProductService) GetAll(cateogryId int) ([]shop.Product, error) {
	return s.repo.GetAll(cateogryId)
}

func (s *ProductService) GetById(categoryId, productId int) (shop.Product, error) {
	return s.repo.GetById(categoryId, productId)
}

func (s *ProductService) Delete(categoryId, productId int) error {
	return s.repo.Delete(categoryId, productId)
}

func (s *ProductService) Update(productId int, input shop.UpdateProductInput) error {
	return s.repo.Update(productId, input)
}
