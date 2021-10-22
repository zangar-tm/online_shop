package service

import (
	"github.com/jmoiron/sqlx"
	shop "github.com/zangar-tm/online_shop"
)

type Authorization interface {
	CreateUser(user shop.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Product interface {
}

type Category interface {
}

type Service struct {
	Authorization
	Product
	Category
}

func NewService(db *sqlx.DB) *Service {
	return &Service{}
}
