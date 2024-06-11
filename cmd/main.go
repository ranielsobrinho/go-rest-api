package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	//repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//usecase
	ProductUsecase := usecase.NewProductUsecase(ProductRepository)

	// controllers
	ProductController := controller.NewProductController(ProductUsecase)

	server.GET("/products", ProductController.GetProducts)

	server.POST("/products", ProductController.CreateProduct)

	server.GET("/products/:productId", ProductController.GetProductById)

	server.DELETE("/products/:productId", ProductController.DeleteProductById)

	server.PUT("/products/:productId", ProductController.UpdateProduct)

	server.Run(":5000")
}
