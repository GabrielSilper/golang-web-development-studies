package controllers

import (
	"golang-web-development-studies/dao/product_dao"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(resp http.ResponseWriter, req *http.Request) {
	products := product_dao.FindAll()
	temp.ExecuteTemplate(resp, "Index", products)
}
