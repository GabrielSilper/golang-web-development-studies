package main

import (
	"database/sql"
	"fmt"
	"golang-web-development-studies/models"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	PORT = 8000
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func connectToDb() *sql.DB {
	uri := "postgresql://root:password@localhost/loja?sslmode=disable"
	db, err := sql.Open("postgres", uri)

	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	http.HandleFunc("/", index)

	fmt.Println("Servidor rodando na porta", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)

	if err != nil {
		log.Fatalln(err)
	}
}

func index(resp http.ResponseWriter, req *http.Request) {
	db := connectToDb()
	defer db.Close()

	products := []models.Product{}

	allProductsQuery, err := db.Query("SELECT * FROM products")

	if err != nil {
		panic(err.Error())
	}

	for allProductsQuery.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err = allProductsQuery.Scan(&id, &name, &description, &price, &quantity)
		if err != nil {
			panic(err.Error())
		}

		product := models.Product{
			Name:        name,
			Description: description,
			Price:       price,
			Quantity:    quantity,
		}

		products = append(products, product)
	}

	temp.ExecuteTemplate(resp, "Index", products)
}
