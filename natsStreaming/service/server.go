package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	model "service/service/models"
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

func stanMsgHandler(cache map[string]string, db *sql.DB) stan.MsgHandler {
	return func(m *stan.Msg) {
		var order model.Order
		//if !(json.Valid(m.Data)) {
		//	fmt.Println("This is a bad json")
		//}
		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			fmt.Println("Given data has wrong format: ", err)
		}

		fmt.Printf("New uid recieved %#v\n", order.OrderUid)
		//fmt.Println(string(m.Data))
		//todo check it's the right data
		if order.OrderUid != "" {
			//I don't check if keys are correct

			cache[order.OrderUid] = string(m.Data)

			q := "INSERT INTO orders (uid, data) VALUES ($1, $2) RETURNING uid"
			_, err = db.Exec(q, order.OrderUid, string(m.Data))
		}
	}
}

func httpHandler(cache map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "404 not found. ", http.StatusNotFound)
			return
		}

		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "form.html")
		case "POST":

			// Call ParseForm() to parse the raw query and update r.PostForm and r.Form.
			fmt.Fprintf(w, "<a href=\"/\">Home page</a>")
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "<div>ParseForm() err: %v</div>", err)
				return
			}
			uid := r.FormValue("order_uid")
			data, ok := cache[uid]
			if ok != true {
				fmt.Fprintf(w, "<div>There is no data with associated order_uid %s<div>", uid)
			}
			var order model.Order
			json.Unmarshal([]byte(data), &order)

			s, _ := json.MarshalIndent(order, "", "\t")
			fmt.Println(string(s))
			fmt.Fprintln(w, "<div>", string(s), "</div>")
			//fmt.Fprintf(w, "Address = %s\n", address)
		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	})
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

// type message stan.MsgHandler
//
//	func (m *message) receiveMsg() message {
//		//fmt.Printf("Received a message: %s\n", string(m.Data))
//
//		//stan.AckWait(stan.DefaultAckWait)
//		var order model.Order //todo
//		err := json.Unmarshal(m.Data, &order)
//		if err != nil {
//			fmt.Println("Given data has wrong format: ", err)
//		}
//		fmt.Printf("%#v\n", order.OrderUid)
//		//fmt.Println(string(m.Data))
//		//todo check it's the right data
//		//todo check uid exist
//		//todo store in cache data
//
//		sqlStatement := "INSERT INTO orders (uid, data) VALUES ($1, $2) RETURNING uid"
//
//		_, err = db.Exec(sqlStatement, order.OrderUid, string(m.Data))
//
//		//fmt.Println("New order added to db: ", res)
//		//if uid != "" {
//		//}
//	}

func restoreCache(db *sql.DB) map[string]string {
	cache := map[string]string{}

	q := "SELECT uid, data FROM orders LIMIT $1"
	rows, err := db.Query(q, 50)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		uid, data := "", ""
		err = rows.Scan(&uid, &data)
		if err != nil {
			fmt.Println(err)
		}

		cache[uid] = data
		fmt.Println("found and saved in cache: ", uid)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}

	return cache
}

func main() {
	sc, err := stan.Connect("test-cluster", "waif")
	if err != nil {
		panic(err)
	}

	db := ConnectDB()
	cache := restoreCache(db)

	sub, _ := sc.Subscribe("orders", stanMsgHandler(cache, db), stan.StartWithLastReceived())

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

	//helloHandler := func(w http.ResponseWriter, req *http.Request) {
	//	io.WriteString(w, "Hello, world!\n")
	//}
	//
	//http.HandleFunc("/hello", helloHandler)
	//log.Fatal(http.ListenAndServe(":8080", nil))

	//server block
	//sample
	fmt.Println(cache)

	http.Handle("/", httpHandler(cache))

	fmt.Printf("Starting HTTP server...\n")
	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatal(err)
	}
}
