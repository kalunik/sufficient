package natsStreaming

import (
	conf "app/config"
	model "app/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	"time"
)

var previousUid string

func ParseMsg(msg []byte) model.Order {
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
		order := ParseMsg(m.Data)
		if !isCorrectUid(order.OrderUid) {
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

func WorkerStanMsg(cache model.Cache, db *sql.DB) {
	for {
		sc, err := stan.Connect(conf.StanCluster, conf.StanClient, stan.NatsURL(conf.StanUrl))
		if err != nil {
			fmt.Println("Connecting stan err: ", err)
		}

		sub, _ := sc.Subscribe(conf.StanSubj, stanMsgHandler(cache, db),
			stan.DurableName(conf.StanDurableName), stan.StartWithLastReceived())

		err = sub.Unsubscribe()
		if err != nil {
			fmt.Println("Unsubscribing STAN chanel err: ", err)
		}
		err = sc.Close()
		if err != nil {
			fmt.Println("Closing STAN connection err: ", err)
		}

		time.Sleep(conf.ReconnectInterval)
	}
}
