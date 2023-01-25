package main

import (
	conf "app/config"
	"app/internal/controller"
	ns "app/pkg/natsStreaming"
	pg "app/pkg/postgres"
	"fmt"
	"net/http"
)

func main() {
	db := pg.ConnectDB()
	cache := pg.RestoreCache(db)

	fmt.Println("STAN will check", conf.StanSubj, "every", conf.ReconnectInterval)
	go ns.WorkerStanMsg(cache, db)

	http.Handle("/", controller.HttpHandler(cache))

	fmt.Println("Starting HTTP server on port", conf.ServPort)
	if err := http.ListenAndServe(conf.ServPort, nil); err != nil {
		fmt.Println("HTTP server err: ", err)
	}

	err := db.Close()
	if err != nil {
		fmt.Println("Closing db err: ", err)
	}
}
