// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
//во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import "fmt"

func main() {
	firstCh := make(chan int)
	secondCh := make(chan int)

	go func() {
		for v := range firstCh {
			secondCh <- v * 2
		}
	}()

	array := [10]int{1, 4, 2, 5, 6, 3, 7, 9, 0, 8}
	for _, v := range array {
		firstCh <- v
		fmt.Println(<-secondCh) // ждем пока функция поместит результат в канал
	}
	close(firstCh)
	close(secondCh)
}
