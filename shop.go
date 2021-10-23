package shop

import "errors"

type Category struct {
	Id    int    `json:"id" db:"id"`
	Title string `json:"title" db:"title" binding:"required"`
}

type CategoryProducts struct {
	id         int
	ProductId  int
	CategoryId int
}

type Product struct {
	Id          int    `json:"id" db:"id"`
	Title       string `json:"title" db:"title" binding:"required"`
	Description string `json:"description" db:"description"`
	Image       string `json:"image" db:"image" binding:"required"`
	Price       int    `json:"price" db:"price" binding:"required"`
}

type UpdateProductInput struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Price       *int    `json:"price"`
	Image       *string `json:"image"`
}

func (i UpdateProductInput) Validate() error {
	if i.Title == nil && i.Description == nil && i.Price == nil && i.Image == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateCategoryInput struct {
	Title *string `json:"title"`
}

func (i UpdateCategoryInput) Validate() error {
	if i.Title == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
