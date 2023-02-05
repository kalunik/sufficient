package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	str, err := unpackString("a4bc2d5e2")
	if err != nil {
		fmt.Println("Unpacking err: ", err)
	}
	fmt.Println(str)
}

func unpackString(str string) (string, error) {
	var err error
	var count bytes.Buffer     // может заменить string.Builder
	var result strings.Builder //отсутсвует метод чтения
	var previousVal rune
	for i, value := range str {
		if unicode.IsLetter(value) {
			err := writeToResult(value, previousVal, &count, &result) // только по указателю
			if err != nil {
				return "", err
			}
			previousVal = value
		}
		if unicode.IsNumber(value) {
			if i == 0 {
				return "", errors.New("Non valid string. It's starts with a digit.")
			}
			count.WriteRune(value) //так как ограничений по количеству знаков числа нету, то мы будем накапливать в буфер
			if i-1 == len(str) {   // если последняя цифра
				err := writeToResult(value, previousVal, &count, &result)
				if err != nil {
					return "", err
				}
			}
			continue
		}
	}
	return result.String(), err
}

func writeToResult(value rune, previousVal rune, count *bytes.Buffer, result *strings.Builder) error {
	if count.Len() > 0 { //пора сгружать подстроку, так как уже новая подстрока готова записываться
		num, err := strconv.Atoi(count.String())
		if err != nil {
			return err
		}
		for i := 1; i < num; i++ { //отсчет с 1, так как первую руну сразу в результат отправляем
			result.WriteString(string(previousVal))
		}
		count.Reset() //сбрасываем буфер, чтобы записывать новое
	}
	if count.Len() == 0 { //наполняем буфер подстроки, проверку вероятно можно будет убрать
		result.WriteRune(value)
	}
	return nil
}
