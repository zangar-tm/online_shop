package service

import (
	shop "github.com/zangar-tm/online_shop"
	"github.com/zangar-tm/online_shop/pkg/repository"
)

type CommentService struct {
	repo repository.Comment
}

func NewCommentService(repo repository.Comment) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) Create(userId, productId int, comment shop.Comment) (int, error) {
	return s.repo.Create(userId, productId, comment)
}

func (s *CommentService) GetAll(productId int) ([]shop.Comment, error) {
	return s.repo.GetAll(productId)
}
