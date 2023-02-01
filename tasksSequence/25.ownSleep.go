// Реализовать собственную функцию sleep.
package main

import (
	"fmt"
	"time"
)

func main() {
	t := 5

	ownSleep(t)
	fmt.Println("`Спали`", t, "секунд")
}

func ownSleep(t int) {
	timer := time.NewTimer(time.Duration(t) * time.Second)
	<-timer.C //ждем пока время выйдет
}
