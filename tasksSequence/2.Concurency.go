//2. Написать программу, которая конкурентно
//рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"os"
	"sync"
)

func powWait(numbers []int) {
	wg := &sync.WaitGroup{}

	for _, v := range numbers { //range итерируется по слайсу
		wg.Add(1)                            //увеличиваем счетчик ожидания на 1
		go func(v int, wg *sync.WaitGroup) { //запускаем анонимную функцию в горутине
			squared := v * v
			fmt.Fprintf(os.Stdout, "%d ", squared)
			wg.Done() //уменьшаем счётчик на 1
		}(v, wg) //неопределенное поведение, если не передать в горутину число
	}
	wg.Wait()
}

func powChan(numbers []int) { //альтернативный вариант с каналами
	squared := make(chan int, 0) //небуферизированный канал

	for _, v := range numbers {
		go func(v int, in chan int) {
			in <- v * v //результат отправляем в канал
		}(v, squared)
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Fprintf(os.Stdout, "%d ", <-squared) //основная горутина ждет, читает из канала 5 раз, проходит дальше
	}
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	fmt.Println("Wait:")
	powWait(numbers)

	fmt.Println("\nChannels:")
	powChan(numbers)
}
