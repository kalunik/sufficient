//1. Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

type Human struct {
	Id   int
	Name string
}

func (h *Human) SetName(val string) {
	h.Name = val
}

func (h *Human) GetName() string {
	return h.Name
}

type Action struct {
	Human
}

func (a *Action) GetName() string { // переопределяем метод Human, для различия добавим к имени "А"
	return "A" + a.Name
}

func main() {
	human := Human{1, "Ivan"}
	action := &Action{human}

	action.SetName("John")
	fmt.Println("Human name -", human.GetName())
	fmt.Println("Action name -", action.GetName())           //приоритет у action getname
	fmt.Println("Action human name", action.Human.GetName()) // getname родителя с помощью полного селектора
}
