package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=postgres dbname=cola_db user=root password=password sslmode=disable")

	if err != nil {
		panic(err)
	}
}
