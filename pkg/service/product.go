package service

import (
	"github.com/zangar-tm/online_shop/pkg/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

// func (s *Service) Create(product shop.Product) (int, error) {
// 	//return s.repo.Create(product)
// }

// func (s *Service) GetAll() ([]shop.Product, error) {

// }

// func (s *Service) GetById(productId int) (shop.Product, error) {

// }

// func (s *Service) Delete(productId int) error {

// }

// func (s *Service) Update(productId int, input shop.UpdateProductInput) error {

// }
