package repository

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/zangar-tm/online_shop/models"
)

type ProductPostgres struct {
	db *sqlx.DB
}

func NewProductPostgres(db *sqlx.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(categoryId int, product models.Product) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var productId int
	createProductQuery := fmt.Sprintf("INSERT INTO %s (title, description, image, price) values ($1, $2, $3, $4) RETURNING id", productsTable)

	row := tx.QueryRow(createProductQuery, product.Title, product.Description, product.Image, product.Price)
	err = row.Scan(&productId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createCategoriesProductsQuery := fmt.Sprintf("INSERT INTO %s (category_id, product_id) values ($1, $2)", categoriesProductsTable)
	_, err = tx.Exec(createCategoriesProductsQuery, categoryId, productId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return productId, tx.Commit()
}

func (r *ProductPostgres) GetAll(categoryId int) ([]models.Product, error) {
	var products []models.Product

	query := fmt.Sprintf("SELECT pt.id, pt.title, pt.price, pt.description, pt.image FROM %s pt INNER JOIN %s cpt on pt.id = cpt.product_id WHERE cpt.category_id = $1", productsTable, categoriesProductsTable)
	err := r.db.Select(&products, query, categoryId)

	return products, err
}
func (r *ProductPostgres) GetById(categoryId, productId int) (models.Product, error) {
	var product models.Product

	query := fmt.Sprintf(`SELECT pt.id, pt.title, pt.price, pt.description, pt.image FROM %s pt 
																INNER JOIN %s cpt on cpt.product_id=pt.id WHERE cpt.category_id=$1 AND cpt.product_id=$2`, productsTable, categoriesProductsTable)
	err := r.db.Get(&product, query, categoryId, productId)

	return product, err
}

func (r *ProductPostgres) Delete(categoryId, productId int) error {
	query := fmt.Sprintf("DELETE FROM %s pt USING %s cpt WHERE pt.id = cpt.product_id AND cpt.category_id=$1 AND cpt.product_id=$2",
		productsTable, categoriesProductsTable)
	_, err := r.db.Exec(query, categoryId, productId)

	return err
}

func (r *ProductPostgres) Update(productId int, input models.UpdateProductInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *input.Title)
		argId++
	}

	if input.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *input.Description)
		argId++
	}

	if input.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, *input.Price)
		argId++
	}

	if input.Image != nil {
		setValues = append(setValues, fmt.Sprintf("image=$%d", argId))
		args = append(args, *input.Image)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s pt SET %s FROM %s cpt WHERE pt.id = cpt.product_id AND cpt.product_id = $5`, productsTable, setQuery, categoriesProductsTable)
	args = append(args, productId)

	_, err := r.db.Exec(query, args...)
	return err
}
