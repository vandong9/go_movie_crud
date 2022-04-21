package routes

import (
	"github.com/gorilla/mux"
	"github.com/vandong9/go_eleven_basic_project/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router mux.Router) {
	router.HandleFunc("/book/")
}