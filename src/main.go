package main

import (
	"fmt"

	"github.com/qadrina/go-crud-sqlserver-basic/src/config"
	"github.com/qadrina/go-crud-sqlserver-basic/src/entities"
	"github.com/qadrina/go-crud-sqlserver-basic/src/models"
)

func main() {
	//Demo1()
	Demo2()
}

func Demo1() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.FindAll()
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("-----------------")
			}
		}
	}
}

func Demo2() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.Search("")
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("-----------------")
			}
		}
	}
}

func Demo3() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		products, err2 := productModel.SearchByPrice(20, 35)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			for _, product := range products {
				fmt.Println(product.ToString())
				fmt.Println("-----------------")
			}
		}
	}
}

func Demo4() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product, err2 := productModel.Find(12)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println(product.ToString())
		}
	}
}

func Demo5() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product := entities.Product{
			Name:     "abc",
			Price:    4.5,
			Quantity: 9,
			Status:   true,
		}
		err2 := productModel.Create(&product)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("Product Info")
			fmt.Println(product.ToString())
		}
	}
}

func Demo6() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product := entities.Product{
			Id:       1,
			Name:     "abc",
			Price:    4.5,
			Quantity: 9,
			Status:   true,
		}
		rowsAffected, err2 := productModel.Update(&product)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("rowsAffected: ", rowsAffected)
			fmt.Println(product.ToString())
		}
	}
}

func Demo7() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		rowsAffected, err2 := productModel.Delete(2)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("rowsAffected: ", rowsAffected)
		}
	}
}

/*
func Demo5() {
	db, err := config.GetDB()
	if err != nil {
		fmt.Println(err)
	} else {
		productModel := models.ProductModel{
			Db: db,
		}
		product := entities.Product{
			Name:     "abc",
			Price:    4.5,
			Quantity: 9,
			Status:   true,
		}
		rowsAffected, err2 := productModel.Create(&product)
		if err2 != nil {
			fmt.Println(err2)
		} else {
			fmt.Println("rowsAffected: ", rowsAffected)
			fmt.Println("Product Info")
			fmt.Println(product.ToString())
		}
	}
}
*/

/*
tutorial from https://www.youtube.com/watch?v=OylDD1rVPz8
*/
