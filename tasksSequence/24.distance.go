// Разработать программу нахождения расстояния между двумя точками,
// которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (first *Point) GetDistance(second *Point) float64 { //можно и не как метод реализовать
	powSubX := math.Pow(float64(second.x-first.x), 2)
	powSubY := math.Pow(float64(second.y-first.y), 2)
	return math.Sqrt(powSubX + powSubY)
}

func main() {
	first := NewPoint(3, 4)
	second := NewPoint(3, 5)

	fmt.Println(first.GetDistance(second))
}
