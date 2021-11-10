package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zangar-tm/online_shop/models"
)

func (h *Handler) createComment(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	productId, err := strconv.Atoi(c.Param("prod_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid product id param")
		return
	}
	var input models.Comment
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Comment.Create(userId, productId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type getAllCommentsResponse struct {
	Data []models.Comment `json:"data"`
}

func (h *Handler) getComments(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("prod_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid product id param")
		return
	}
	comments, err := h.services.Comment.GetAll(productId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllCommentsResponse{
		Data: comments,
	})
}
