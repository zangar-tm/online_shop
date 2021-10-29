package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	shop "github.com/zangar-tm/online_shop"
)

type CartPostgres struct {
	db *sqlx.DB
}

func NewCartPostgres(db *sqlx.DB) *CartPostgres {
	return &CartPostgres{db: db}
}

func (r *CartPostgres) Create(userId int, input shop.UsersCart) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var id int
	addToCartQuery := fmt.Sprintf("INSERT INTO %s (user_id, product_id) VALUES ($1, $2) RETURNING id", cartTable)
	row := tx.QueryRow(addToCartQuery, userId, input.ProductId)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *CartPostgres) GetAll(userId int) ([]shop.MyCart, error) {
	var mycart []shop.MyCart

	query := fmt.Sprintf("SELECT pt.id, pt.title, pt.price FROM %s pt INNER JOIN %s ct on pt.id = ct.product_id WHERE ct.user_id = $1", productsTable, cartTable)
	err := r.db.Select(&mycart, query, userId)

	return mycart, err
}

func (r *CartPostgres) GetById(productId int) (shop.Product, error) {
	var product shop.Product

	query := fmt.Sprintf(`SELECT id, title, price, description, image FROM %s WHERE id=$1`, productsTable)
	err := r.db.Get(&product, query, productId)

	return product, err
}
func (r *CartPostgres) Delete(productId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", cartTable)

	_, err := r.db.Exec(query, productId)

	return err
}
