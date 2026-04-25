package models	

import
(
	"github.com/jinzhu/gorm"
	"github.com/parthbhardwaj93/Bookstore-Management-APIs/package/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
