package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/stan.go"
	"io"
	"os"
)

func main() {
	jsonPtr := flag.String("json", "../model.json", "use a custom json")
	flag.Parse()

	jsonFile, err := os.Open(*jsonPtr)
	if err != nil {
		fmt.Println(err)
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(jsonFile)

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	//byteValue = []byte("Qwerty") //todo test it

	sc, err := stan.Connect("test-cluster", "john")
	if err != nil {
		panic(err)
	}

	err = sc.Publish("orders", []byte(byteValue))
	if err != nil {
		return
	}
}
