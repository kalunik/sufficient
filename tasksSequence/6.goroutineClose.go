// Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("info:", "closing goroutine:", "closing channel")
	{
		ch := make(chan string)
		go func() {
			for {
				v, ok := <-ch
				if !ok { //завершаем горутину при закрытии канала
					fmt.Println("Finish")
					return
				}
				fmt.Println(v)
			}
		}()

		ch <- "Hey"
		ch <- "guys"
		close(ch)               //закрытие канала
		time.Sleep(time.Second) //дадим немного времени на закрытие горутины
	}
	fmt.Println("\ninfo:", "closing goroutine:", "when 'select' triggers on `done` chanel")
	{
		ch := make(chan string)
		done := make(chan struct{})
		go func(ch chan string, done chan struct{}) { //sender
			for {
				select { //когда что-то будет в канале done, горутина завершится
				case ch <- "foo":
				case <-done:
					close(ch)
					return
				}
				time.Sleep(1 * time.Second)
			}
		}(ch, done)

		go func() {
			time.Sleep(3 * time.Second) //завершим через 3 секунды
			done <- struct{}{}          //не займет места
		}()

		for v := range ch { // range заканчивает работу, когда канал закрыт. _, ok := ch && ok == false
			fmt.Println("Read: ", v)
		}

		fmt.Println("Finish")
	}
	fmt.Println("\ninfo:", "closing goroutine:", "using context cancelFunc")
	{
		ch := make(chan struct{})
		ctx, cancelFunc := context.WithCancel(context.Background())

		go func(ctx context.Context) {
			for {
				select {
				case <-ctx.Done(): //получим после 3 секунд, когда сработает cancel context
					ch <- struct{}{} //теперь основная горутина может продолжить
					return
				default:
					fmt.Println("foo...")
				}

				time.Sleep(1 * time.Second)
			}
		}(ctx)

		go func() {
			time.Sleep(3 * time.Second)
			cancelFunc() // теперь есть что прочитать из ctx.Done
		}()

		<-ch
		fmt.Println("Finish")
	}
}
