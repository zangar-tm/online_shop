package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zangar-tm/online_shop/models"
)

func (h *Handler) addToCart(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var input models.UsersCart
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Cart.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})

}

type getMyCartResponse struct {
	Data []models.MyCart `json:"data"`
}

func (h *Handler) getUsersProducts(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	mycart, err := h.services.Cart.GetAll(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, getMyCartResponse{
		Data: mycart,
	})
}

func (h *Handler) getUsersProductById(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("prod_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid product id param")
		return
	}

	product, err := h.services.Cart.GetById(productId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) deleteUsersProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("prod_id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid product id param")
		return
	}

	err = h.services.Cart.Delete(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}
