package router

import (
	"GoProductLatihan/handler"
	"GoProductLatihan/modules/product"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine{

	productRepository := product.NewRepository()

	productService := product.NewService(productRepository)

	productHandler := handler.NewProductHandler(productService)

	router := gin.Default()

	api := router.Group("/api/v1")

	api.GET("/products", productHandler.GetProducts)
	api.POST("/products", productHandler.CreateProduct)
	api.GET("/products/:id", productHandler.GetProduct)
	api.PUT("/products/:id", productHandler.UpdateProduct)
	api.DELETE("/products/:id", productHandler.DeleteProduct)
	return router
}