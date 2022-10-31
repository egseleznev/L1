// Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»)
// Символы могут быть unicode

package main

import "fmt"

func reverseStr(str string) string {

	if len(str) < 1 { // если строка пустая, то возвращаем ее
		return ""
	}

	runes := []rune(str) // приводим строку к слайсу рун

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 { // идем по i слева направо, по j справа налево
		runes[i], runes[j] = runes[j], runes[i] // меняем буквы местами
	}

	return string(runes) // приводим слайс рун к строке и возвращаем ее
}

func main() {
	str := "строка которую перевернем"

	valid := "менревереп юуроток акортс"

	fmt.Println(reverseStr(str) == valid)
}
