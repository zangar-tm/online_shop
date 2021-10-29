package service

import (
	shop "github.com/zangar-tm/online_shop"
	"github.com/zangar-tm/online_shop/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

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

type Comment interface {
	Create(userId, productId int, comment shop.Comment) (int, error)
	GetAll(productId int) ([]shop.Comment, error)
}

type Cart interface {
	Create(userId int, input shop.UsersCart) (int, error)
	GetAll(userId int) ([]shop.MyCart, error)
	GetById(productId int) (shop.Product, error)
	Delete(productId int) error
}

type Service struct {
	Authorization
	Product
	Category
	Comment
	Cart
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Category:      NewCategoryService(repo.Category),
		Product:       NewProductService(repo.Product),
		Comment:       NewCommentService(repo.Comment),
		Cart:          NewCartService(repo.Cart),
	}
}
