package controllers

import (
	"golang-web-development-studies/dao/product_dao"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(resp http.ResponseWriter, req *http.Request) {
	products := product_dao.FindAll()
	temp.ExecuteTemplate(resp, "Index", products)
}

func New(resp http.ResponseWriter, req *http.Request) {
	temp.ExecuteTemplate(resp, "New", nil)
}

func Insert(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		name := req.FormValue("nome")
		description := req.FormValue("descricao")
		price, err := strconv.ParseFloat(req.FormValue("preco"), 64)
		if err != nil {
			log.Println("Erro ao converter o preço, erro:", err.Error())
		}

		quantity, err := strconv.Atoi(req.FormValue("quantidade"))
		if err != nil {
			log.Println("Erro ao converter a quantidade, erro:", err.Error())
		}

		err = product_dao.Create(name, description, price, quantity)
		if err != nil {
			log.Println("Erro ao criar o produto, erro:", err.Error())
		} else {
			log.Println("Produto", name, "criado com sucesso!")
		}
	}
	http.Redirect(resp, req, "/", http.StatusMovedPermanently)
}
