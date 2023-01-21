package main

import (
	"database/sql"
	"fmt"
	"github.com/nats-io/stan.go"
	"time"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "wjonatho"
	pass   = "My8es1P4ss"
	dbname = "stan"
)

func ConnectDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close() //todo

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func main() {

	//db := ConnectDB()

	sc, err := stan.Connect("test-cluster", "waif")
	if err != nil {
		panic(err)
	}
	time.Sleep(10)
	// Simple Async Subscriber
	sub, _ := sc.Subscribe("orders", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}, stan.DeliverAllAvailable())

	//delivered, err := sub.Delivered()
	//if err != nil {
	//	fmt.Println("Can't check delivered")
	//} else {
	//	fmt.Println("How many is delivered?", delivered)
	//}

	//var order model.Order		//todo
	//err = json.Unmarshal(byteValue, &order)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(order)

	//var result map[string]interface{}
	//json.Unmarshal([]byte(byteValue), &result)

	// Unsubscribe
	err = sub.Unsubscribe()
	if err != nil {
		fmt.Println("Can't unsubscribe")
	}

	// Close connection
	err = sc.Close()
	if err != nil {
		fmt.Println("Can't close connection")
	}
}
