package main

import (
	"flag"
	"fmt"
	"github.com/nats-io/stan.go"
	"io"
	"os"
)

func main() {
	jsonPtr := flag.String("json", "order1.json", "use a custom json")
	flag.Parse()

	jsonFile, err := os.Open(*jsonPtr)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}

	sc, err := stan.Connect("test-cluster", "john")
	if err != nil {
		panic(err)
	}
	defer sc.Close()

	err = sc.Publish("orders", []byte(byteValue))
	if err != nil {
		fmt.Println("Can't publish: ", err)
	}
	fmt.Println(*jsonPtr, "was sent")
}
