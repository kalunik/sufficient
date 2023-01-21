package main

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"io"
	"os"
)

func main() {
	jsonFile, err := os.Open("../model.json")
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
	fmt.Println(byteValue)

	sc, err := stan.Connect("test-cluster", "john")
	if err != nil {
		panic(err)
	}

	err = sc.Publish("orders", []byte(byteValue))
	if err != nil {
		return
	}
}
