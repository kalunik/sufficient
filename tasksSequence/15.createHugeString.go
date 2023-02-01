//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//
//var justString string
//func someFunc() {
//v := createHugeString(1 << 10)
//justString = v[:100]
//}
//
//
//func main() {
//someFunc()
//}

package main

import (
	"fmt"
)

// глобальную переменную лучше не использовать, тогда её не придется хранить на протяжении всей программы

// Если символ занимает больше 1 байта, то символов при выводе будет, меньше чем 100,
// и их число будет зависеть от размера символа
// Лучше будет использовать возвращаемый тип слайс рун, тогда мы сможем взять 100 рун и вывести их
func createHugeString(size int) []rune {
	var hugeString []rune

	for i := 0; i < size; i++ { //размер - 2^10
		hugeString = append(hugeString, 'Ⓕ') //весит 3 байта, строка была бы небольшой
	}
	return hugeString
}

func someFunc() []rune {
	v := createHugeString(1 << 10)
	return v[:100]
}

func main() {
	justString := string(someFunc())
	fmt.Println(len(justString), justString) //итоговый размер этой строки из 100 символов 300 байт
}
