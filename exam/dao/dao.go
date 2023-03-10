package dao

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:123456@(localhost:3306)/exam?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		log.Fatal(err)
		return
	}
	DB = db
	fmt.Println(DB.Ping())
}
