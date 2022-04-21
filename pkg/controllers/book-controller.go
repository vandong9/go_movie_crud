package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/vandong9/go_eleven_basic_project/pkg/utils"
	"github.com/vandong9/go_eleven_basic_project/pkg/models"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBook:=models.GetAllBooks()
}