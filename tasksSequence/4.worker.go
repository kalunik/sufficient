//4. Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func getPoolSize() int {
	var workerPoolSize int

	fmt.Println("Enter number of workers:")
	if _, err := fmt.Scan(&workerPoolSize); err != nil {
		fmt.Println("Scan err:", err)
	}
	return workerPoolSize
}

func writeToChan(workCh chan int, cancelFunc context.CancelFunc) {
	cancelCh := make(chan os.Signal, 1) //буфферизированный канал, чтобы не пропустить сигнал, если мы не готовы, а сигнал послан
	go func(cancelCh chan os.Signal) {
		signal.Notify(cancelCh, os.Interrupt) //ожидает SIGINT, отправит в канал
	}(cancelCh)

LOOP: //маркер цикла
	for i := 0; ; i++ {
		select { //мултиплексор
		case <-cancelCh: //при получении сигнала сработает
			cancelFunc() //при вызове функци отправится в канал ctx.Done
			break LOOP
		default: // выполняется если нет других кейсов, которые готовы
			workCh <- i
		}
	}
	close(workCh) //канал закрывается на стороне отправителя, из него еще можно читать
}

func readWorker(in chan int, ctx context.Context, workerSerial int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done(): //как только отпработает CancelFunc
			fmt.Fprintln(os.Stdout, "worker #", workerSerial, ": I've done ")
			wg.Done() //подсчет завершившихся воркеров
			return
		case num, ok := <-in:
			if ok {
				fmt.Fprintln(os.Stdout, "worker #", workerSerial, "output:", num)
			}
		}
	}
}

func main() {

	workerPoolSize := getPoolSize()

	ctx, cancelFunc := context.WithCancel(context.Background()) //для передачи в воркер информации о завершении, при получении сигнала
	wg := &sync.WaitGroup{}
	wg.Add(workerPoolSize)

	workCh := make(chan int)
	for i := 0; i < workerPoolSize; i++ {
		go readWorker(workCh, ctx, i, wg)
	}

	writeToChan(workCh, cancelFunc)
	wg.Wait() //ждем завершения воркеров
}
