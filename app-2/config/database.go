package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var MasterDB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "stnduser:stnduser@tcp(192.168.1.145:6033)/go_crud?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database Master")
	MasterDB = db
}
