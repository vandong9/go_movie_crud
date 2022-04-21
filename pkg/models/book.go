package models

import(
	"github.com/jinzhu/gorm"
	"github.com/vandong9/go_eleven_basic_project@go_movie_crud/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publicatio string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}