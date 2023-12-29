package models

import (
	"database/sql"

	"github.com/qadrina/go-crud-sqlserver-basic/src/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() ([]entities.Product, error) {
	rows, err := productModel.Db.Query("select * from product")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			var status bool
			err2 := rows.Scan(&id, &name, &price, &quantity, &status)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
					Status:   status,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) Search(keyword string) ([]entities.Product, error) {
	rows, err := productModel.Db.Query("select * from product where name like ?", "%"+keyword+"%")
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			var status bool
			err2 := rows.Scan(&id, &name, &price, &quantity, &status)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
					Status:   status,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) SearchByPrice(min, max float64) ([]entities.Product, error) {
	rows, err := productModel.Db.Query("select * from product where price between ? and ?", min, max)
	if err != nil {
		return nil, err
	} else {
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			var status bool
			err2 := rows.Scan(&id, &name, &price, &quantity, &status)
			if err2 != nil {
				return nil, err2
			} else {
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
					Status:   status,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) Find(id int64) (entities.Product, error) {
	rows, err := productModel.Db.Query("select * from product where id = ?", id)
	if err != nil {
		return entities.Product{}, err
	} else {
		var product entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			var status bool
			err2 := rows.Scan(&id, &name, &price, &quantity, &status)
			if err2 != nil {
				return entities.Product{}, err2
			} else {
				product = entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
					Status:   status,
				}

			}
		}
		return product, nil
	}
}

func (productModel ProductModel) Create(product *entities.Product) error {
	row, err := productModel.Db.Query("insert into product values(?,?,?,?) select convert(bigint, SCOPE_IDENTITY());", product.Name, product.Price, product.Quantity, product.Status)
	if err != nil {
		return err
	} else {
		var newId int64
		row.Next()
		row.Scan(&newId)
		product.Id = newId
		return nil
	}
}

func (productModel ProductModel) Update(product *entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("update product set name = ?, price = ?, quantity = ?, status = ? where id = ?", product.Name, product.Price, product.Quantity, product.Status, product.Id)
	if err != nil {
		return 0, err
	} else {

		return result.RowsAffected()
	}
}

func (productModel ProductModel) Delete(id int64) (int64, error) {
	result, err := productModel.Db.Exec("delete from product  where id = ?", id)
	if err != nil {
		return 0, err
	} else {

		return result.RowsAffected()
	}
}

/*
func (productModel ProductModel) Create(product *entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("insert into product values(?,?,?,?)", product.Name, product.Price, product.Quantity, product.Status)
	if err != nil {
		return 0, err
	} else {
		product.Id, _ = result.LastInsertId()
		rowsAffected, _ := result.RowsAffected()
		return rowsAffected, nil
	}
}
*/
