package repository

import (
	"github.com/jmoiron/sqlx"
	shop "github.com/zangar-tm/online_shop"
)

type Authorization interface {
	CreateUser(user shop.User) (int, error)
	GetUser(username, password string) (shop.User, error)
}

type Product interface {
	Create(product shop.Product) (int, error)
	GetAll() ([]shop.Product, error)
	GetById(productId int) (shop.Product, error)
	Delete(productId int) error
	Update(productId int, input shop.UpdateProductInput) error
}

type Category interface {
	Create(list shop.Category) (int, error)
	GetAll() ([]shop.Category, error)
	GetById(listId int) (shop.Category, error)
	Delete(listId int) error
	// Update(listId int, input shop.UpdateCategoryInput) error
}

type Repository struct {
	Authorization
	Product
	Category
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Category: NewCategoryRepository(db),
	}
}
