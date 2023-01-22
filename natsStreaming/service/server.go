package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	model "service/service/models"
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

//func receiveMsg(db *sql.DB, m *stan.Msg) stan.MsgHandler {
//	return func(msg *stan.Msg) {
//		//fmt.Printf("Received a message: %s\n", string(m.Data))
//
//		var order model.Order //todo
//		err := json.Unmarshal(m.Data, &order)
//		if err != nil {
//			fmt.Println("Given data has wrong format: ", err)
//		}
//		fmt.Printf("%#v\n", order.OrderUid)
//		fmt.Println(string(m.Data))
//		//todo check it's the right data
//		//todo check uid exist
//		//create table if not exist
//
//		//todo write to db (u_id | jsonb)
//		uid := ""
//		sqlStatement := `
//		INSERT INTO orders (uid, data)
//		VALUES ($1, $2)`
//		err = db.QueryRow(sqlStatement, order.OrderUid, string(m.Data)).Scan(&uid)
//		if err != nil {
//			panic(err)
//		}
//		if uid != "" {
//			fmt.Println("New order added to db: ", uid)
//		}
//	}
//}

func main() {

	db := ConnectDB()

	sc, err := stan.Connect("test-cluster", "waif")
	if err != nil {
		panic(err)
	}
	time.Sleep(10)
	// Simple Async Subscriber
	//var m *stan.Msg
	sub, _ := sc.Subscribe("orders", func(m *stan.Msg) {
		//fmt.Printf("Received a message: %s\n", string(m.Data))

		//stan.AckWait(stan.DefaultAckWait)
		var order model.Order //todo
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			fmt.Println("Given data has wrong format: ", err)
		}
		fmt.Printf("%#v\n", order.OrderUid)
		//fmt.Println(string(m.Data))
		//todo check it's the right data
		//todo check uid exist
		//todo store in cache data

		sqlStatement := "INSERT INTO orders (uid, data) VALUES ($1, $2) RETURNING uid"

		res, err := db.Exec(sqlStatement, order.OrderUid, string(m.Data))

		fmt.Println("New order added to db: ", res)
		//if uid != "" {
		//}
	}, stan.StartWithLastReceived())

	//delivered, err := sub.Delivered()
	//if err != nil {
	//	fmt.Println("Can't check delivered")
	//} else {
	//	fmt.Println("How many is delivered?", delivered)
	//}

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
