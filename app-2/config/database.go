package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var MasterDB *sql.DB

func ConnectDB() {
	db, err := sql.Open("mysql", "crud:usercrud@tcp(192.168.50.28:3306)/test.go-crud?parseTime=true")
	if err != nil {
		panic(err)
	}

	log.Println("Connected to database Master")
	MasterDB = db
}
