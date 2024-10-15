package controller

// Aqui criamos o controle que vai receber a requisição
// essa parte serve tanto para receber a requisição como para tratar ela e apresentar uma response
import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//criamos a estrutura para o controle

type productController struct {
	productUseCase  usecase.ProductUsecase
}



func NewProductController( usecase usecase.ProductUsecase) productController {

	return productController{
			productUseCase: usecase,
	}

}
//Função the FATO que vai tratar a requisição 
func (p *productController) GetProducts(ctx *gin.Context){

	products, err := p.productUseCase.GetProducts()
	if (err != nil) {

		ctx.JSON(http.StatusInternalServerError, err)
	}
	
	ctx.JSON(http.StatusOK,products)
}


func(p *productController) CreateProduct(ctx *gin.Context) {

	var product model.Product 
	err := ctx.BindJSON(&product)

	if err != nil{
		ctx.JSON(http.StatusBadRequest, err )
		return 
	}

	insertedProduct , err := p.productUseCase.CreateProduct(product)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, err )
		return 
	}


	ctx.JSON(http.StatusCreated, insertedProduct)

}
func (p *productController) GetProductbyId(ctx *gin.Context){

	id := ctx.Param("productId")

	if id == ""{
		response := model.Response{
			Message: "Id do produto não pode ser nulo",
		}
		ctx.JSON(http.StatusBadRequest,response)
		return
	}

	productId, err := strconv.Atoi(id)
	
	if err != nil {
		response := model.Response{
			Message: "Id do produto precisa ser um número",
		}
		ctx.JSON(http.StatusBadRequest,response)
		return

	}


	product, err := p.productUseCase.GetProductbyId(productId)
	if (err != nil) {

		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {

		response := model.Response{
			Message: "Produto não foi encontrado na base de dados ",
		}
		ctx.JSON(http.StatusNotFound,response)
		return

	}
	
	ctx.JSON(http.StatusOK,product)
}

