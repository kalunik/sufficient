// Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package main

import (
	"fmt"
)

func sorting(ints []int, low int, high int) {
	if low < high {

		pivot := ints[low] //выбор опорного эл-та
		i := low + 1

		for j := low; j <= high; j++ {
			if ints[j] < pivot {
				ints[i], ints[j] = ints[j], ints[i]
				i++
			}
		}

		ints[low], ints[i-1] = ints[i-1], pivot

		sorting(ints, low, i-2)
		sorting(ints, i, high)
	}
}

func quicksort(ints []int) {
	if len(ints) >= 2 { //незачем сортировать уже готовое
		sorting(ints, 0, len(ints)-1)
	}
}

func main() {
	ints := []int{124, 15, 62, 65622, 32, -32, 34, 51, -25, 0}
	fmt.Println(ints, "before")
	quicksort(ints)
	fmt.Println(ints, "after")
}
