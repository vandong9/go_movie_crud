package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vandong9/go_eleven_basic_project@go_movie_crud/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routers.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}