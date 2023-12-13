package config

import(
	"github.com/jinzhu/gorm"
	//"github.com/jinzhu/gorm/dialects/mysql"
)

var(
	db * gorm.DB
)

func Connect(){
	d, err := gorm.Open("mysql","root:1234@tcp(127.0.0.1:3306)/go_bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database")
	}	
	db =d
}


func GetDB() *gorm.DB{
	return db
}