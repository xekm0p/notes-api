package utility

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	dsn := "admin:admin@tcp(127.0.0.1:3309)/note_service"
	db, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	log.Println("DB Connection established...")
	return db
}
