package main

import (
	"fmt"
	"golang-web-development-studies/routes"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	PORT = 8000
)

func main() {
	routes.LoadRoutes()

	fmt.Println("Servidor rodando na porta", PORT)
	err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil)

	if err != nil {
		log.Fatalln(err)
	}
}
