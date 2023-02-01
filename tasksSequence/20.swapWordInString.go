// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».
package main

import (
	"fmt"
	"strings"
)

func swapWords(s string) string {
	splited := strings.Split(s, " ") //делаем слайс
	reversed := ""
	for _, word := range splited {
		reversed = word + " " + reversed //собираем в обратном порядке
	}
	return reversed
}

func main() {
	s := "person snow dog sun public"

	fmt.Println(s, "–", swapWords(s))
}
