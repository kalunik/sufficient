// Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"

func main() {
	first := []string{"apple", "banana", "orange"}
	second := []string{"microsoft", "apple", "jetbrains"}
	common := append(first, second...) //всё в кучу

	inter := make(map[string]int)
	for _, v := range common {
		_, ok := inter[v] //уже было это слово?
		if !ok {
			inter[v] = 1
		} else {
			inter[v]++ //то что надо, мно-во неуникальных значений
		}
	}

	result := []string{}
	for k, v := range inter {
		if v > 1 {
			result = append(result, k)
		}
	}
	fmt.Println(result)
}
