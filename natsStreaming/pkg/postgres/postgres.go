package postgres

import (
	conf "app/config"
	model "app/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	pgsqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Pass, conf.Dbname)

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

func RestoreCache(db *sql.DB) model.Cache {

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
