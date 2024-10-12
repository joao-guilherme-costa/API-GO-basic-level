package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}



func NewProductRepository (connection *sql.DB) ProductRepository{
	return ProductRepository {
		connection: connection,
	}
}

func (pr *ProductRepository ) GetProducts() ([]model.Product, error){

	querry := "SELECT id, product_name, price FROM product"
	rows , err := pr.connection.Query(querry)
	
	if err != nil{
		fmt.Println(err)
		return []model.Product{}, err
		 
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next(){
		err = rows.Scan(
			&productObj.ID,
			&productObj.Price,
			&productObj.Name,)

			if err != nil {
				fmt.Println(err)
				return []model.Product{}, err
			}

			productList = append(productList, productObj)


	}
	rows.Close()

	return productList , nil

}