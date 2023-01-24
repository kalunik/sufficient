package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	model "service/service/models"
	"time"
)

var previousUid string

func parseMsg(msg []byte) model.Order {
	var order model.Order

	err := json.Unmarshal(msg, &order)
	if err != nil {
		order.OrderUid = ""
	}
	return order
}

func isCorrectUid(uid string) bool {
	if uid == "" || uid == previousUid {
		return false
	}
	previousUid = uid
	return true
}

func stanMsgHandler(cache model.Cache, db *sql.DB) stan.MsgHandler {
	return func(m *stan.Msg) {
		order := parseMsg(m.Data)
		if !isCorrectUid(order.OrderUid) {
			fmt.Println("STAN listen ...")
			return
		}

		fmt.Println("STAN got NEW order: ", order.OrderUid)
		cache.Set(order.OrderUid, string(m.Data))

		q := `INSERT INTO orders (uid, data) VALUES ($1, $2) RETURNING uid`
		_, err := db.Exec(q, order.OrderUid, string(m.Data))
		if err != nil {
			fmt.Println("Insert query err: ", err)
		}

	}
}

func workerStanMsg(cache model.Cache, db *sql.DB) {
	for {
		sc, err := stan.Connect(stanCluster, stanClient)
		if err != nil {
			fmt.Println("Connecting stan err: ", err)
		}

		sub, _ := sc.Subscribe(stanSubj, stanMsgHandler(cache, db),
			stan.DurableName(stanDurableName), stan.StartWithLastReceived())

		err = sub.Unsubscribe()
		if err != nil {
			fmt.Println("Unsubscribing STAN chanel err: ", err)
		}
		err = sc.Close()
		if err != nil {
			fmt.Println("Closing STAN connection err: ", err)
		}

		time.Sleep(intervalMsgCheck)
	}
}
