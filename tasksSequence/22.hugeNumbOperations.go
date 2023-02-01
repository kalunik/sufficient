// Разработать программу, которая перемножает, делит, складывает,
// вычитает две числовых переменных a,b, значение которых > 2^20.
package main

import (
	"fmt"
	"math/big"
)

func main() {
	n1 := "2097155"
	n2 := "1048578"
	fmt.Println(n1, n2)

	c := new(big.Int)
	a, ok := new(big.Int).SetString(n1, 10) //bigInt из строки, с указанной системой
	b, oki := new(big.Int).SetString(n2, 10)
	if ok == false || oki == false {
		fmt.Println("Error args")
	}

	fmt.Println("Сложение: ", c.Add(a, b))
	fmt.Println("Вычитание: ", c.Sub(a, b))
	fmt.Println("Деление: ", c.Div(a, b))
	fmt.Println("Умножение: ", c.Mul(a, b))
}
