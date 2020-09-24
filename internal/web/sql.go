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

func titlesusersfromdb(db *sql.DB) []portaldata {
	var pd []portaldata
	str := "SELECT name,title FROM users;"
	rows, err := db.Query(str)
	if err != nil {
		log.Printf("Query")
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var title string
		if err := rows.Scan(&name,&title); err != nil {
			log.Printf("scan")
			log.Fatal(err)
		}
		pd = append(pd, portaldata{name,title})
	}
	return pd
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

