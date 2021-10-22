package shop

import "errors"

type Category struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

type CategoryProducts struct {
	id         int
	ProductId  int
	CategoryId int
}

type Product struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Image       string `json:"image" db:"image"`
	Price       int    `json:"price" db:"price"`
}

type UpdateProductInput struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	Price       *int    `json:"price"`
	Image       *string `json:"image"`
}

func (i UpdateProductInput) Validate() error {
	if i.Name == nil && i.Description == nil && i.Price == nil && i.Image == nil {
		return errors.New("update structure has no values")
	}

	return nil
}

type UpdateCategoryInput struct {
	Name *string `json:"name"`
}

func (i UpdateCategoryInput) Validate() error {
	if i.Name == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
