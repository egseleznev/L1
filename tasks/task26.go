// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc)
// Функция проверки должна быть регистронезависимой

package main

import (
	"fmt"
	"strings"
)

func checkUnique(str string) bool {

	uniqueSymbols := make(map[rune]struct{}) // множество для проверки уникальности символа

	for _, v := range strings.ToLower(str) { // проходимся по каждому символу строки, приведя ее к нижнем регистру

		if _, ok := uniqueSymbols[v]; ok { // проверяем символ на наличие в множестве
			return false // если есть, возвращаем false, так как есть повторяющиеся элементы

		} else {
			uniqueSymbols[v] = struct{}{} // если нет, добавляем символ в множество
		}
	}

	return true // возвращаем true, так как все символы уникальные
}

func main() {
	str := "abCdefAaf"

	fmt.Println(checkUnique(str) == false)
}
