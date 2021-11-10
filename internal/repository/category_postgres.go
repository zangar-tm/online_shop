package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zangar-tm/online_shop/models"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}
func (r *CategoryPostgres) Create(category models.Category) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	createCategoryQuery := fmt.Sprintf("INSERT INTO %s (title) VALUES ($1) RETURNING id", categoriesTable)
	row := tx.QueryRow(createCategoryQuery, category.Title)
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}
	return id, tx.Commit()
}

func (r *CategoryPostgres) GetAll() ([]models.Category, error) {
	var categories []models.Category

	query := fmt.Sprintf("SELECT id, title FROM %s", categoriesTable)
	err := r.db.Select(&categories, query)

	return categories, err
}

func (r *CategoryPostgres) GetById(categoryId int) (models.Category, error) {
	var category models.Category

	query := fmt.Sprintf(`SELECT id, title FROM %s WHERE id=$1`, categoriesTable)
	err := r.db.Get(&category, query, categoryId)

	return category, err
}

func (r *CategoryPostgres) Delete(categoryId int) error {
	query := fmt.Sprintf("DELETE FROM %s  WHERE id =$1", categoriesTable)
	_, err := r.db.Exec(query, categoryId)

	return err
}

func (r *CategoryPostgres) Update(categoryId int, input models.UpdateCategoryInput) error {
	var arg string

	if input.Title != nil {
		arg = *input.Title
	}

	query := fmt.Sprintf("UPDATE %s ct SET title=$1 WHERE ct.id = $2", categoriesTable)
	_, err := r.db.Exec(query, arg, categoryId)
	return err
}
