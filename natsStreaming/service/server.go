package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	model "service/service/models"
	"time"
)

const (
	stanCluster      = "test-cluster"
	stanClient       = "waif"
	stanSubj         = "orders"
	stanDurableName  = "john"
	intervalMsgCheck = 15 * time.Second
	servPort         = ":8888"
)

func prepareData(uid string, cache model.Cache, w http.ResponseWriter) string {
	data, ok := cache.Get(uid)
	if ok != true {
		fmt.Fprintf(w, "<div>There is no data with associated order_uid %s<div>", uid)
	}
	//var order model.Order
	//json.Unmarshal([]byte(data), &order)
	//
	//s, _ := json.MarshalIndent(order, "\n", "\t")
	return data
}

func httpHandler(cache model.Cache) http.Handler {
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
			orderJson := prepareData(uid, cache, w)
			fmt.Fprintln(w, "<div>", orderJson, "</div>")

		default:
			fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
		}
	})
}

func main() {
	db := connectDB()
	cache := restoreCache(db)

	fmt.Println("STAN will check", stanSubj, "every", intervalMsgCheck)
	go workerStanMsg(cache, db)

	http.Handle("/", httpHandler(cache))

	fmt.Println("Starting HTTP server on port", servPort)
	if err := http.ListenAndServe(servPort, nil); err != nil {
		fmt.Println("HTTP server err: ", err)
	}

	err := db.Close()
	if err != nil {
		fmt.Println("Closing db err: ", err)
	}
}
