package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	shop "github.com/zangar-tm/online_shop"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}
func (r *CategoryPostgres) Create(category shop.Category) (int, error) {
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

func (r *CategoryPostgres) GetAll() ([]shop.Category, error) {
	var categories []shop.Category

	query := fmt.Sprintf("SELECT id, title FROM %s", categoriesTable)
	err := r.db.Select(&categories, query)

	return categories, err
}

func (r *CategoryPostgres) GetById(categoryId int) (shop.Category, error) {
	var category shop.Category

	query := fmt.Sprintf(`SELECT id, title FROM %s WHERE id=$1`, categoriesTable)
	err := r.db.Get(&category, query, categoryId)

	return category, err
}

func (r *CategoryPostgres) Delete(categoryId int) error {
	query := fmt.Sprintf("DELETE FROM %s  WHERE id =$1", categoriesTable)
	_, err := r.db.Exec(query, categoryId)

	return err
}

// func (r *CategoryPostgres) Update(categoryId int, input shop.UpdateCategoryInput) error {
// 	var setValue string
// 	args := make([]interface{}, 0)
// 	argId := 1

// 	if input.Title != nil {
// 		setValue = append(setValue, fmt.Sprintf("title=$%d", argId))
// 		args = append(args, *input.Title)
// 	}

// 	query := fmt.Sprintf("UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id=$%d AND ul.user_id=$%d",
// 		categoriesTable, setQuery, argId, argId+1)
// 	args = append(args, listId, userId)

// 	logrus.Debugf("updateQuery: %s", query)
// 	logrus.Debugf("args: %s", args)

// 	_, err := r.db.Exec(query, args...)
// 	return err
// }
