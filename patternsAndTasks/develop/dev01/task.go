package main

/*
=== Базовая задача ===

Создать программу печатающую точное время с использованием NTP библиотеки.Инициализировать как go module.
Использовать библиотеку https://github.com/beevik/ntp.
Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Программа должна быть оформлена с использованием как go module.
Программа должна корректно обрабатывать ошибки библиотеки: распечатывать их в STDERR и возвращать ненулевой код выхода в OS.
Программа должна проходить проверки go vet и golint.
*/

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
)

func getNTPTime(host string) error {
	time, err := ntp.Time(host)
	if err != nil {
		return err
	}
	fmt.Println(time)
	return nil
}

func main() {
	err := getNTPTime("0.beevik-ntp.pool.ntp.org")
	if err != nil { //вынес обработку ошибки, так удобнее тестировать
		fmt.Println("NTP error :", err) //использует Fprintln(os.Stdout, )
		os.Exit(1)                      // в случае ошибки сразу выйдем
	}

}
