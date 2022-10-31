// Удалить i-ый элемент из слайса

package main

import (
	"fmt"
	"reflect"
)

func removeElement(slice []int, element int) []int {

	if len(slice) < 1 {
		return []int{} // если слайс пустой, нечего удалять, выходим
	}

	return append(slice[:element], slice[element+1:]...) // с помощью срезов берем левую часть до элемента, который нужно удалить, и добавляем правую часть от этого элемента, таким образом необходимый элемент удаляется из слайса
}

func main() {
	slice := []int{5, 2, 1, 3, 4, 8, 7, 15, 12, 16, 23, 25, 27, 30, 1, 3, 6, 4}

	valid := []int{5, 2, 1, 4, 8, 7, 15, 12, 16, 23, 25, 27, 30, 1, 3, 6, 4}

	slice = removeElement(slice, 3)

	fmt.Println(reflect.DeepEqual(slice, valid))
}
