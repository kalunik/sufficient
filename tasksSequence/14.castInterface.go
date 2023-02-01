// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.
package main

import "fmt"

func typeDefine(value interface{}) string {
	switch value.(type) { // я хотел сначала в свитче кейсы отлавлиать `case value.(int)`, но так не вышло, зато есть `type`
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan int:
		return "chan int"
	case chan string:
		return "chan string"
	default:
		return "UNKNOWN"
	}
}

func main() {
	intCh := make(chan int)
	stringCh := make(chan string)
	fmt.Println(typeDefine(50), typeDefine("HELL"),
		typeDefine(false), typeDefine(intCh), typeDefine(stringCh))
}
