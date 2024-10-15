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
	//camada de repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	//camada usecase
	ProductUseCase := usecase.NewProductUsecase(ProductRepository)
	
	//camada de controles 
	ProductController := controller.NewProductController(ProductUseCase)

	
	server.GET("/products",ProductController.GetProducts)
	server.POST("/product",ProductController.CreateProduct)
	server.GET("/product/:productId",ProductController.GetProductbyId)
	//criar rotas de put e delete e adicionar autenticação jwt 
	server.Run(":8080")
	
}




