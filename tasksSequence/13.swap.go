// Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main() {
	first := 14
	second := 50

	fmt.Println("1st-", first, "2nd-", second)
	first, second = second, first
	fmt.Println("1st-", first, "2nd-", second)
}
