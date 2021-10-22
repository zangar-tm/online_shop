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
	Create(userId int, list shop.Product) (int, error)
	GetAll(userId int) ([]shop.Product, error)
	GetById(userId, listId int) (shop.Product, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input shop.UpdateProductInput) error
}

type Category interface {
	Create(userId int, list shop.Category) (int, error)
	GetAll(userId int) ([]shop.Category, error)
	GetById(userId, listId int) (shop.Category, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input shop.UpdateCategoryInput) error
}

type Repository struct {
	Authorization
	Product
	Category
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
