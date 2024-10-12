package controller

// Aqui criamos o controle que vai receber a requisição
// essa parte serve tanto para receber a requisição como para tratar ela e apresentar uma response
import (
	"go-api/usecase"
	"net/http"

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