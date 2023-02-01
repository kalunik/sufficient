// Реализовать паттерн «адаптер» на любом примере.
package main

import "fmt"

type computer interface {
	usingAPFS()
}

type client struct {
}

func (c *client) usingAPFS(com computer) {
	com.usingAPFS()
}

type macOS struct {
}

func (m *macOS) usingAPFS() {
	fmt.Println("Using APFS")
}

type windows struct {
}

func (w *windows) usingNTFS() {
	fmt.Println("Using NTFS")
}

type windowsAdapter struct {
	pc *windows
}

func (w *windowsAdapter) usingAPFS() {
	w.pc.usingNTFS()
}

func main() {
	client := &client{}

	mac := &macOS{}
	client.usingAPFS(mac)

	myPC := &windows{} //клиент не может использовать apfs на myPC, так как такого метода не существует
	pcAdapter := &windowsAdapter{
		pc: myPC,
	} //поэтому необходим адаптер
	client.usingAPFS(pcAdapter)
}
