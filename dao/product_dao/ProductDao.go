package product_dao

import (
	"golang-web-development-studies/db"
	"golang-web-development-studies/models"
)

func FindAll() []models.Product {
	dbInstance := db.ConnectToDb()
	defer dbInstance.Close()

	products := []models.Product{}

	rows, err := dbInstance.Query("SELECT * FROM products")
	if err != nil {
		panic(err.Error())
	}

	for rows.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = rows.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product := models.NewProduct(id, quantity, name, description, price)
		products = append(products, product)
	}

	return products
}

func Create(name, description string, price float64, quantity int) error {
	dbInstance := db.ConnectToDb()
	defer dbInstance.Close()

	insertQuery, err := dbInstance.Prepare("INSERT INTO products (name, description, price, quantity) values ($1, $2, $3, $4)")

	if err != nil {
		return err
	}

	_, err = insertQuery.Exec(name, description, price, quantity)

	if err != nil {
		return err
	}

	return nil
}
