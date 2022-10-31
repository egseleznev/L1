// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество

package main

import "fmt"

// Почти то же самое, что и хеш-таблица, но без сохранения значения
// Используем если нам нужна только быстрая проверка вхождения
// В качестве значения нужно лишь указать пустое значение struct{}, занимающее 0 байт
func createStringSet(sequence []string) map[string]struct{} {

	set := make(map[string]struct{})

	for _, v := range sequence {
		set[v] = struct{}{}
	}
	return set
}

func main() {

	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(createStringSet(sequence))
}
