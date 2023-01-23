package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"log"
	"net/http"
	model "service/service/models"
)

const (
	stanCluster = "test-cluster"
	stanClient  = "waif"
	stanSubj    = "orders"
	servPort    = ":8888"
)

func prepareData(uid string, cache map[string]string, w http.ResponseWriter) string {
	data, ok := cache[uid]
	if ok != true {
		fmt.Fprintf(w, "<div>There is no data with associated order_uid %s<div>", uid)
	}
	var order model.Order
	json.Unmarshal([]byte(data), &order)

	s, _ := json.MarshalIndent(order, "\n", "\t")
	return string(s)
}

func httpHandler(cache map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Path != "/" {
			http.Error(w, "404 not found. ", http.StatusNotFound)
			return
		}

		switch r.Method {
		case "GET":
			http.ServeFile(w, r, "./static/form.html")
		case "POST":

			fmt.Fprintf(w, "<a href=\"/\">Home page</a>")

			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "<div>ParseForm() err: %v</div>", err)
				return
			}
			uid := r.FormValue("order_uid")
			s := prepareData(uid, cache, w)
			fmt.Fprintln(w, "<div>", s, "</div>")

		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	})
}

func main() {
	sc, err := stan.Connect(stanCluster, stanClient)
	if err != nil {
		panic(err)
	}

	db := connectDB()
	cache := restoreCache(db)

	sub, _ := sc.Subscribe(stanSubj, stanMsgHandler(cache, db),
		stan.StartWithLastReceived())

	err = sub.Unsubscribe()
	if err != nil {
		fmt.Println("Unsubscribing STAN chanel err: ", err)
	}
	err = sc.Close()
	if err != nil {
		fmt.Println("Closing STAN connection err: ", err)
	}
	err = db.Close()
	if err != nil {
		fmt.Println("Closing db err: ", err)
	}

	http.Handle("/", httpHandler(cache))

	fmt.Printf("Starting HTTP server...\n")
	if err := http.ListenAndServe(servPort, nil); err != nil {
		log.Fatal(err)
	}
}
