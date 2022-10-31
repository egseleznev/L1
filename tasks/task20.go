// Разработать программу, которая переворачивает слова в строке
// Пример: «snow dog sun — sun dog snow

package main

import (
	"fmt"
	"strings"
)

func reverseWords(str string) string {

	if len(str) < 1 { // если строка пустая, то возвращаем ее
		return ""
	}

	words := strings.Split(str, " ") // разбиваем строку на слова

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 { // идем по i слева направо, по j справа налево
		words[i], words[j] = words[j], words[i] // меняем буквы местами
	}

	return strings.Join(words, " ") // склеиваем слова в строку
}

func main() {
	str := "строка слова которой перевернем"

	valid := "перевернем которой слова строка"

	fmt.Println(reverseWords(str) == valid)
}
