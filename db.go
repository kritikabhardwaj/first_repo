package db

import "github.com/go-pg/pg"

//ConnectDB to db
func ConnectDB() (db *pg.DB) {
	db = pg.Connect(&pg.Options{
		User:     "root",
		Password: "tolexo",
		Database: "postgres",
		Addr:     "localhost:5432",
	})
	return db
}
