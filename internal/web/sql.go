package web

import (
	"database/sql"
	"fmt"
	"log"
)

func QueryServer(db *sql.DB, ch chan) {
	select {
		msg := <-ch
		row := db.Query(msg)
		row.Scan()
	}
}

func CreateDataBase() chan {
	sqlchan := make(chan string)
	connStr := fmt.Sprintf("user=%s dbname=%s",config.SQLUser, config.SQLDatabaseName)
	db, err := sql.Open("postgres",connStr)

	if err != nil {
		log.Fatal(err)
	}

	return sqlchan
}

