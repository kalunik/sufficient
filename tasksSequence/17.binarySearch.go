// Реализовать бинарный поиск встроенными методами языка.
package main

import (
	"fmt"
	"sort"
)

func binarySearch(ints []int, item int) int {
	left := 0 // границы поиска
	right := len(ints) - 1

	for left <= right { // пока границы не сократились до одного
		mid := left + right // проверяем центр
		sus := ints[mid]
		if sus == item { // если нашли возвращаем индекс
			return mid
		}
		if sus > item { // двигаем границу к центру
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	ints := []int{124, 15, 62, 65622, 32, -32, 34, 51, -25, 0}
	var find int

	sort.Ints(ints) // необходимо отсортировать предварительно
	fmt.Println("Enter number:")
	fmt.Scan(&find)
	index := binarySearch(ints, find)
	if index != -1 {
		fmt.Println("Found number -", ints[index], "| position - ", index)
	} else {
		fmt.Println("Not found")

	}
}
