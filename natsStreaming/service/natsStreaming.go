package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/nats-io/stan.go"
	model "service/service/models"
)

func stanMsgHandler(cache map[string]string, db *sql.DB) stan.MsgHandler {
	return func(m *stan.Msg) {

		var order model.Order

		err := json.Unmarshal(m.Data, &order)
		if err != nil {
			fmt.Println("Given data has wrong format: ", err)
		}

		if order.OrderUid != "" {
			cache[order.OrderUid] = string(m.Data)

			q := `INSERT INTO orders (uid, data) VALUES ($1, $2) RETURNING uid`
			_, err = db.Exec(q, order.OrderUid, string(m.Data))
			if err != nil {
				fmt.Println("Insert query err: ", err)
			}
		}
	}
}
