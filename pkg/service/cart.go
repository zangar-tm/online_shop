package service

import (
	shop "github.com/zangar-tm/online_shop"
	"github.com/zangar-tm/online_shop/pkg/repository"
)

type CartService struct {
	repo repository.Cart
}

func NewCartService(repo repository.Cart) *CartService {
	return &CartService{repo: repo}
}

func (s *CartService) Create(userId int, input shop.UsersCart) (int, error) {
	return s.repo.Create(userId, input)
}

func (s *CartService) GetAll(userId int) ([]shop.MyCart, error) {
	return s.repo.GetAll(userId)
}
func (s *CartService) GetById(productId int) (shop.Product, error) {
	return s.repo.GetById(productId)
}
func (s *CartService) Delete(productId int) error {
	return s.repo.Delete(productId)
}
