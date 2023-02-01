//Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
	"time"
)

type concMap struct {
	sync.RWMutex
	numbers map[int]int
}

func newMap() *concMap {
	return &concMap{numbers: make(map[int]int)}
}

func (r *concMap) Store(key int, value int) {
	r.Lock() //сможет залочиться, если есть лок на чтение
	r.numbers[key] = value
	r.Unlock()
}

func (r *concMap) Load(key int) (int, bool) {
	r.RLock()
	defer r.RUnlock()
	value, ok := r.numbers[key]
	return value, ok
}

func main() {

	myMap := newMap()

	for work := 0; work < 4; work++ {
		go func() {
			for i := 0; ; i++ {
				myMap.Store(i, i%4)
			}

		}()
	}
	time.Sleep(3 * time.Second)
	val1, _ := myMap.Load(1)
	val2, _ := myMap.Load(10)
	fmt.Println(val1, val2)
}
