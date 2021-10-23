package repository

import (
	"github.com/jmoiron/sqlx"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

// func (r *ProductPostgres) Create(product shop.Product) (int, error) {

// }
