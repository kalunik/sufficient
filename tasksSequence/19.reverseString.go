// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»).
// Символы могут быть unicode.
package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseString(s string) string {
	runes := []rune(s)
	reversed := make([]rune, 0)                           //руны при создании ждут конкретного len
	for i := utf8.RuneCountInString(s) - 1; i >= 0; i-- { //количество рун в строке
		reversed = append(reversed, runes[i])
	}
	return string(reversed)
}

func main() {
	s := "главрыба😀"
	res := reverseString(s)
	fmt.Println(res)
}
