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
