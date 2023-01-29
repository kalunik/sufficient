//4. Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"os"
	"os/signal"
)

func worker(in <-chan int, quit chan os.Signal, workerSerial int) {
	//LOOP:
	for {
		select {
		case <-in:
			fmt.Fprintln(os.Stdout, "worker #", workerSerial, "output:", <-in)
			//case <-quit:
			//	fmt.Fprintln(os.Stdout, "worker #", workerSerial, ": I've done ")
			//	break LOOP
		}
	}
}

func main() {
	// Set up channel on which to send signal notifications.
	// We must use a buffered channel or risk missing the signal
	// if we're not ready to receive when the signal is sent.
	c := make(chan os.Signal, 1)

	// Passing no signals to Notify means that
	// all signals will be sent to the channel.
	signal.Notify(c)

	// Block until any signal is received.
	s := <-c
	fmt.Println("Got signal:", s)

	//var workersQty int
	//
	//fmt.Println("Enter number of workers:")
	//if _, err := fmt.Scan(&workersQty); err != nil {
	//	fmt.Println("Scan err:", err)
	//	return
	//}
	//
	//workCh := make(chan int)
	//cancelCh := make(chan os.Signal, 1)
	//for i := 0; i < workersQty; i++ {
	//	go worker(workCh, cancelCh, i)
	//}
	//
	//go func(cancelCh chan os.Signal) {
	//	signal.Notify(cancelCh, os.Interrupt)
	//}(cancelCh)
	//
	//for i := 0; ; i++ {
	//	select {
	//	case <-cancelCh:
	//		close(workCh)
	//	default:
	//		workCh <- i
	//
	//	}
	//}

}
