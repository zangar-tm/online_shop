package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zangar-tm/online_shop/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		categories := api.Group("/categories")
		{
			categories.POST("/", h.createCategory)
			categories.GET("/", h.getCategories)
			categories.GET("/:id", h.getCategoryById)
			categories.PUT("/:id", h.updateCategory)
			categories.DELETE("/:id", h.deleteCategory)

			products := categories.Group(":id/products")
			{
				products.POST("/", h.createProduct)
				products.GET("/", h.getProducts)
				products.GET("/:prod_id", h.getProductById)
				products.PUT("/:prod_id", h.updateProduct)
				products.DELETE("/:prod_id", h.deleteProduct)
			}
		}
	}
	return router
}
