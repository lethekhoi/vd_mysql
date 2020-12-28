package models

import (
	"database/sql"
	"errors"
	"fmt"
	"vd_mysql/entities"
)

type ProductModel struct {
	Db *sql.DB
}

func (productModel ProductModel) FindAll() (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	{
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			}
			{
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		return products, nil
	}
}

func (productModel ProductModel) Search(keyword string) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product where name like ?", "%"+keyword+"%")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	{
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			}
			{
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}

		}
		return products, nil
	}
}
func (productModel ProductModel) SearchPrice(min, max float64) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product where price >= ? and price <= ?", min, max)
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	{
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			}
			{
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}
		}
		err = rows.Err()
		if err != nil {
			return nil, err
		}
		return products, nil
	}
}
func (productModel ProductModel) SearchID(ID int64) (product []entities.Product, err error) {
	rows, err := productModel.Db.Query("select * from product where id = ?", ID)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	{
		var products []entities.Product
		for rows.Next() {
			var id int64
			var name string
			var price float64
			var quantity int64
			err2 := rows.Scan(&id, &name, &price, &quantity)
			if err2 != nil {
				return nil, err2
			}
			{
				product := entities.Product{
					Id:       id,
					Name:     name,
					Price:    price,
					Quantity: quantity,
				}
				products = append(products, product)
			}

		}
		return products, nil
	}
}

func (productModel ProductModel) SearchIDRow(ID int64) (product entities.Product, err error) {
	err = productModel.Db.QueryRow("select * from product where id = ?", ID).Scan(&product.Id, &product.Name, &product.Price, &product.Quantity)
	if err != nil {
		if err == sql.ErrNoRows {
			return product, sql.ErrNoRows
		}
		{
			return product, err
		}
	}
	return product, nil
}

func (productModel ProductModel) Create(product *entities.Product) (err error) {
	result, err := productModel.Db.Exec("insert into product(name, price,quantity) values(?,?,?)", product.Name, product.Price, product.Quantity)

	if err != nil {
		return err
	}
	{
		product.Id, _ = result.LastInsertId()
		return nil
	}
}
func (productModel ProductModel) CreateWithPrepared(product *entities.Product) (err error) {
	fmt.Println("prepare")
	tx, err := productModel.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("insert into product(name, price,quantity) values(?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close() // danger!

	result, err := stmt.Exec(product.Name, product.Price, product.Quantity)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}
	{
		product.Id, _ = result.LastInsertId()
		return nil
	}

}

func (productModel ProductModel) Update(product *entities.Product) (int64, error) {
	result, err := productModel.Db.Exec("update product set name=?, price=?, quantity=? where id=?", product.Name, product.Price,
		product.Quantity, product.Id)

	if err != nil {
		return 0, err
	}

	return result.RowsAffected()

}

func (productModel ProductModel) Delete(id int64) (int64, error) {
	result, err := productModel.Db.Exec("delete from product where id=?", id)

	if err != nil {
		return 0, err
	}
	rowEffected, _ := result.RowsAffected()

	if rowEffected == 0 {
		return 0, errors.New("No data to delete")
	}

	return result.RowsAffected()

}

func (productModel ProductModel) CountProduct() (int64, error) {
	row, err := productModel.Db.Query("select count(*) as count_product from product ")
	if err != nil {
		return 0, err

	}
	{
		var count_product int64
		for row.Next() {
			row.Scan(&count_product)
		}
		return count_product, nil
	}
}
