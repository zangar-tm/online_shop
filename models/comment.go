package models

type Comment struct {
	Id      int    `json:"id" db:"id"`
	Title   string `json:"title" db:"title" binding:"required"`
	Body    string `json:"body" db:"body" binding:"required"`
	User_id int    `json:"user_id" db:"user_id"`
}

type ProductComments struct {
	id        int
	ProductId int
	CommentId int
}
