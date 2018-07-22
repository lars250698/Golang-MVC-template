package main

import (
	"github.com/gorilla/mux"
	"MVC/routes"
	"net/http"
)

var router *mux.Router

func main() {
	router = mux.NewRouter()
	routes.Routes(router)
	http.ListenAndServe(":8080", router)
}
