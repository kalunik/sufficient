// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type counter struct {
	atomic.Uint64
	operationsCount uint64
}

func newCounter() *counter {
	return &counter{}
}

func (c *counter) inc() {
	c.operationsCount = c.Uint64.Add(1) //увеличиваем счётчик на 1
}

func main() {
	total := newCounter()
	for i := 0; i < 1000; i++ {
		go total.inc()
	}
	time.Sleep(2 * time.Millisecond)
	fmt.Println("total operation = ", total.operationsCount)
}
