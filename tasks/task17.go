// Реализовать бинарный поиск встроенными методами языка

package main

import (
	"fmt"
	"sort"
)

// В уже отсортированном массиве берем центральный элемент. Если его значение больше значения искомого элемента, то двигаем границу справа на 1, если меньше - двигаем левую границу
// Таким образом, не приходится проходить по всему массиву целиком, как при линейном поиске со сложностью O(n)
// У бинарного же поиска, сложность в лучшем случае равна O(1), в худшем - O(log n)
// Соответственно, при большой длине получаем выигрыш в производительности относительно линейного метода
func binarySearch(array []int, searchElement int) int {

	if len(array) < 1 {
		return -1 // если массив пустой, возвращаем - 1
	}

	leftBoarder, rightBoarder := 0, len(array)-1 // инициализируем границы

	for leftBoarder <= rightBoarder { // двигаемся пока границы не сойдутся

		middleElementIndex := (leftBoarder + rightBoarder) / 2 // получаем индекс центрального элемента

		if array[middleElementIndex] < searchElement {
			leftBoarder = middleElementIndex + 1 // двигаем границу вправо
		} else if array[middleElementIndex] > searchElement {
			rightBoarder = middleElementIndex - 1 // двигаем границу влево
		} else {
			return middleElementIndex // если величины совпали, элемент найден, возвращаем его индекс
		}
	}

	return -1 // если элемент не найден, возвращаем -1
}

func main() {
	array := []int{5, 2, 1, 3, 4, 8, 7, 15, 12, 16, 23, 25, 27, 30, 1, 3, 6, 4}
	sort.Ints(array)

	fmt.Println(binarySearch(array, 1) == 1)
	fmt.Println(binarySearch(array, 7) == 9)
	fmt.Println(binarySearch(array, 25) == 15)
	fmt.Println(binarySearch(array, 31) == -1)
	fmt.Println(binarySearch(array, 9) == -1)
}
