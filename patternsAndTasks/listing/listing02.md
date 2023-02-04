Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

defer'ы выполняются по типу Last In First Out, а
их аргументы  вычисляются на момент их объявления
return без аргументов автоматически возвращает 
текущее состояние каждого именованного возвращаемого значения

```
