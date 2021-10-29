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
	Create(categoryId int, product shop.Product) (int, error)
	GetAll(categoryId int) ([]shop.Product, error)
	GetById(categoryId, productId int) (shop.Product, error)
	Delete(categoryId, productId int) error
	Update(productId int, input shop.UpdateProductInput) error
}

type Category interface {
	Create(list shop.Category) (int, error)
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

type Repository struct {
	Authorization
	Product
	Category
	Comment
	Cart
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Product:       NewProductPostgres(db),
		Category:      NewCategoryPostgres(db),
		Comment:       NewCommentPostgres(db),
		Cart:          NewCartPostgres(db),
	}
}
