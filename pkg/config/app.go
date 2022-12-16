package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db * gorm.DB
)

func Connect(){
	dsn := "host=127.0.0.1 user=postgres password=olim123 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	fmt.Println("db: connected")
	db=database
}

func GetDB()*gorm.DB{
	return db
}