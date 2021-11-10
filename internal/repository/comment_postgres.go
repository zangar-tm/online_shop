package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/zangar-tm/online_shop/models"
)

type CommentPostgres struct {
	db *sqlx.DB
}

func NewCommentPostgres(db *sqlx.DB) *CommentPostgres {
	return &CommentPostgres{db: db}
}

func (r *CommentPostgres) Create(userId, productId int, comment models.Comment) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var commentId int
	createCommentQuery := fmt.Sprintf("INSERT INTO %s (title, body, user_id) values ($1, $2, $3) RETURNING id", commentsTable)

	row := tx.QueryRow(createCommentQuery, comment.Title, comment.Body, userId)
	err = row.Scan(&commentId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	createProdCommentsQuery := fmt.Sprintf("INSERT INTO %s (product_id, comment_id) values ($1, $2)", productsCommentsTable)
	_, err = tx.Exec(createProdCommentsQuery, productId, commentId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return commentId, tx.Commit()
}

func (r *CommentPostgres) GetAll(productId int) ([]models.Comment, error) {
	var comments []models.Comment
	query := fmt.Sprintf(`SELECT ct.id, ct.title, ct.body, ct.user_id FROM %s ct INNER JOIN %s pct on pct.comment_id = ct.id WHERE pct.product_id = $1`,
		commentsTable, productsCommentsTable)
	if err := r.db.Select(&comments, query, productId); err != nil {
		return nil, err
	}

	return comments, nil
}
