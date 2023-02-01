// Разработать программу, которая проверяет, что все символы в строке
// уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
//
// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false
package main

import (
	"fmt"
	"strings"
)

func uniqueSimbols(str string) bool {
	strSlice := strings.Split(str, "") //разделяем строку на слайс
	unique := make(map[string]bool)
	for _, v := range strSlice {
		_, ok := unique[v]
		if ok == false {
			unique[v] = true //каждую букву заносим в мапу, если попадется повторно - 0
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("abcd", uniqueSimbols("abcd"))
	fmt.Println("abCdefAaf", uniqueSimbols("abCdefAaf"))
	fmt.Println("aabcd", uniqueSimbols("aabcd"))
}
