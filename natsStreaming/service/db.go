package main

import (
	"database/sql"
	"fmt"
	model "service/service/models"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "wjonatho"
	pass   = "My8es1P4ss"
	dbname = "stan"
)

func connectDB() *sql.DB {
	pgsqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)

	db, err := sql.Open("postgres", pgsqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func restoreCache(db *sql.DB) model.Cache {

	q := `SELECT uid, data FROM orders LIMIT $1`
	rows, err := db.Query(q, 50)
	if err != nil {
		fmt.Println("Query err: ", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			fmt.Println("Closing rows err: ", err)
		}
	}(rows)

	cache := &model.InMemoryCache{}
	for rows.Next() {
		uid, data := "", ""
		err = rows.Scan(&uid, &data)
		if err != nil {
			fmt.Println("Scanning rows err: ", err)
		}
		cache.Set(uid, data)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println("Err on rows: ", err)
	}

	return cache
}
