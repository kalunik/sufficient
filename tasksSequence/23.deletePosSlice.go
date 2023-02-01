// Удалить i-ый элемент из слайса.

package main

import "fmt"

func deletePosFromSlice(nums []int, i int) []int {
	return append(nums[:i], nums[i+1:]...) //заносим всё кроме искомого индекса
}

func main() {
	nums := []int{124, 15, 62, 65622, 32, -32, 34, 51, -25, 0}
	fmt.Println("Before", nums)

	deletePosition := 9
	nums = deletePosFromSlice(nums, deletePosition)
	fmt.Println("After", nums)
}
