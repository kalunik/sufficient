//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
	"strconv"
)

func setBit(num *int64, position int, bit int) {
	if bit == 1 {
		*num |= 1 << int64(position) //(3)100000 + 000010
	} else {
		mask := ^(1 << int64(position)) //(1)000010 -> 111101
		*num &= int64(mask)             //(2)100010 * 111101
	}
}

func main() {
	var num int64 = 34

	fmt.Println(strconv.FormatInt(num, 2), "before bit clear")
	setBit(&num, 1, 0)
	fmt.Println(strconv.FormatInt(num, 2), "after bit clear")

	fmt.Println(strconv.FormatInt(num, 2), "before bit set")
	setBit(&num, 1, 1)
	fmt.Println(strconv.FormatInt(num, 2), "after bit set")

}
