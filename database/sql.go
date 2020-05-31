package database

import (
	"log"

	_ "github.com/go-sql-driver/mysql" // mysql dialect - core
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialect for gorm
)

// ConnectSQL to mySQL database
func ConnectSQL() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Parkour3run@/ewallet?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Panic("databases/sql.go : Connection to mySQL is failed ", err)
	}
	log.Println("Connected to mySQL ... succeeded")
	return db
}
