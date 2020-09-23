package web

import (
	"comp-webserver/config"

	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
)

func InsertUser(db *sql.DB, name string, pass string, title string) error {
	str := `INSERT INTO users(name,pass,title) 
	VALUES($1,$2,$3);`
	_, err := db.Exec(str,name,pass,title)
	return err
}

func CreateDB() *sql.DB {
	connStr := fmt.Sprintf("user=%s dbname=%s sslmode=disable",config.SQLUser, config.SQLDatabaseName)

	db, err := sql.Open("postgres",connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

