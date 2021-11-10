package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/zangar-tm/online_shop/models"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(username, password string) (models.User, error)
}

type Product interface {
	Create(categoryId int, product models.Product) (int, error)
	GetAll(categoryId int) ([]models.Product, error)
	GetById(categoryId, productId int) (models.Product, error)
	Delete(categoryId, productId int) error
	Update(productId int, input models.UpdateProductInput) error
}

type Category interface {
	Create(list models.Category) (int, error)
	GetAll() ([]models.Category, error)
	GetById(categoryId int) (models.Category, error)
	Delete(categoryId int) error
	Update(categoryId int, input models.UpdateCategoryInput) error
}

type Comment interface {
	Create(userId, productId int, comment models.Comment) (int, error)
	GetAll(productId int) ([]models.Comment, error)
}

type Cart interface {
	Create(userId int, input models.UsersCart) (int, error)
	GetAll(userId int) ([]models.MyCart, error)
	GetById(productId int) (models.Product, error)
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
