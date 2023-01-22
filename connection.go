package Configs

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


var (
	Dbconn *gorm.DB
)

func DBInit() *gorm.DB {

	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()

	dsn := "host=localhost user=user password=password dbname=my_database port=5432 sslmode=disable TimeZone=Asia/Taipei"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("koneksi kedatabase berhasil")
	Dbconn = db
	//return Dbconn
	return Dbconn

}
