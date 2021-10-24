package service

import (
	shop "github.com/zangar-tm/online_shop"
	"github.com/zangar-tm/online_shop/pkg/repository"
)

type Authorization interface {
	CreateUser(user shop.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Product interface {
	Create(categoryId int, product shop.Product) (int, error)
	GetAll(cateogryId int) ([]shop.Product, error)
	GetById(categoryId, productId int) (shop.Product, error)
	Delete(categoryId, productId int) error
	Update(productId int, input shop.UpdateProductInput) error
}

type Category interface {
	Create(category shop.Category) (int, error)
	GetAll() ([]shop.Category, error)
	GetById(categoryId int) (shop.Category, error)
	Delete(categoryId int) error
	Update(categoryId int, input shop.UpdateCategoryInput) error
}

type Service struct {
	Authorization
	Product
	Category
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Category: NewCategoryService(repo.Category),
		Product:  NewProductService(repo.Product),
	}
}
