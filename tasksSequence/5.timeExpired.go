// Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	timePtr := flag.Int("limit", 2, "use time limit in seconds")
	flag.Parse()

	numbChan := make(chan int)
	timer := time.NewTimer(time.Duration(*timePtr) * time.Second)

	go func(in chan int) {
		for {
			fmt.Println(<-in)
		}
	}(numbChan)

writer:
	for i := 0; ; i++ {
		select {
		case <-timer.C: //канал таймера, сигнализирующи об истечении времени
			fmt.Println("Time's up")
			close(numbChan)
			break writer
		default:
			numbChan <- i
		}
	}
}
