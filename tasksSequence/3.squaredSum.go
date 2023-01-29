//3. Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"os"
	"sync"
)

func powSum(numbers []int) int {
	var sum int
	wg := &sync.WaitGroup{}

	for _, v := range numbers { //range итерируется по слайсу
		wg.Add(1)                                      //увеличиваем счетчик ожидания на 1
		go func(v int, sum *int, wg *sync.WaitGroup) { //wg should not be copied, so pointer
			*sum += v * v
			wg.Done() //уменьшаем счётчик на 1
		}(v, &sum, wg)
	}
	wg.Wait()
	return sum
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}

	sum := powSum(numbers)
	fmt.Fprintf(os.Stdout, "%d ", sum)
}
