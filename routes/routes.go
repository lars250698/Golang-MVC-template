package routes

import (
	"github.com/gorilla/mux"
	"MVC/controllers"
)

func Routes(r *mux.Router) {
	r.HandleFunc("/", controllers.Index)
}
