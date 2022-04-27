package routes

import (
	"github.com/gorilla/mux"
	"github.com/vandong9/go_movie_crud/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.GetBookByID).Methods("DELETE")
}