package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

//estrutura usada para executar queries qsl
type ProductRepository struct {
	connection *sql.DB //o uso de ponteiro é  para manter uma unica conexão reutilizável
}


// é um construtor, recebe a conexão com o banco e retorna uma instância de Productrepository permitindo que crie objetos de forma simples e reutiliza a conexão com o banco de dados 
func NewProductRepository (connection *sql.DB) ProductRepository{
	return ProductRepository {
		connection: connection,
	}
}

func (pr *ProductRepository ) GetProducts() ([]model.Product, error){

	querry := "SELECT id, product_name, price FROM product"
	rows, err := pr.connection.Query(querry)
	
	if err != nil{
		fmt.Println(err)
		return []model.Product{}, err
		 
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next(){
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

			if err != nil {
				fmt.Println(err)
				return []model.Product{}, err
			}

			productList = append(productList, productObj)


	}
	rows.Close()

	return productList , nil
}

func (pr *ProductRepository) CreateProduct(product model.Product) (int,error){
	var id int 
	query , err := pr.connection.Prepare("INSERT INTO product" + 
		"(product_name, price)"	+
		" VALUES ($1 ,$2) RETURNING id")
	if err != nil{
		fmt.Println(err)
		return 0,err 
	}
	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil{
		fmt.Println(err)
		return 0,err 
	}
	query.Close()
	return id , nil
}

func (pr *ProductRepository) GetProductbyId(id_product int) (*model.Product, error){
	
		query , err := pr.connection.Prepare("SELECT * FROM product WHERE id = $1")
		if err != nil{
			fmt.Println(err)
			return nil ,err
		}
		var produto model.Product 

		err = query.QueryRow(id_product).Scan(
			&produto.ID,
			&produto.Name,
			&produto.Price,
		)

		if err != nil{
			if(err == sql.ErrNoRows){
				return nil, nil
			}
				return nil, err
		}
	query.Close()
	return &produto, nil
}		