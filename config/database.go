package config

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"gorm.io/driver/mysql"
)

var (
	host     = "localhost"
	port     = 8889
	user     = "root"
	password = "root"
	dbname   = "simple_chat"
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
		user, password, host, port, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	return db
}
