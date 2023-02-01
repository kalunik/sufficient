// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func main() {
	many := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println("before", many)

	dict := make(map[string]int)
	unique := []string{}

	for _, val := range many {
		dict[val] += 1 //мы теперь знаем кол-во вхождений слова исходное мн-во
		if dict[val] == 1 {
			unique = append(unique, val)
		}
	}
	fmt.Println("after", unique)
}
