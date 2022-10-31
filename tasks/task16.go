// Реализовать быструю сортировку массива (quicksort) встроенными методами языка

package main

import (
	"fmt"
	"reflect"
	"sort"
)

func partition(array []int, leftBorder, rightBorder int) int {

	middleElement := array[(leftBorder+rightBorder)/2] // берем центральный элемент (можно брать также случайный или последний)

	lastLeft, firstRight := leftBorder, rightBorder

	// необходимо найти последний элемент слева, больший чем центральный и первый элемент справа, меньший чем центральный и поменять их местами
	// повторяем процедуру до тех пор, пока индекс левого элемента не станет равным или большим чем индекс правого элемента

	for lastLeft <= firstRight {
		for array[lastLeft] < middleElement {
			lastLeft++
		}

		for array[firstRight] > middleElement {
			firstRight--
		}

		if lastLeft >= firstRight {
			break
		}

		array[lastLeft], array[firstRight] = array[firstRight], array[lastLeft]

		lastLeft++
		firstRight--
	}
	return firstRight // получаем новую границу
}

// Уровень сложности алгоритма O(n log n)
func quickSort(array []int, leftBorder, rightBorder int) {

	if len(array) < 1 { // если массив пустой, выходим
		return
	}

	if leftBorder < rightBorder { // так как функция вызывается рекурсивно, алгоритм завершает работу в случае, когда левая граница становится больше или равной правой
		newBorder := partition(array, leftBorder, rightBorder) // меняем местами элементы двух частей и получаем новую границу
		quickSort(array, leftBorder, newBorder)                // проходим по левой части разделенного массива
		quickSort(array, newBorder+1, rightBorder)             // проходим по правой части разделенного массива
	}
}

func main() {
	array := []int{5, 2, 1, 3, 4, 8, 7, 15, 12, 16, 23, 25, 27, 30, 1, 3, 6, 4}

	valid := array
	sort.Ints(valid)

	quickSort(array, 0, len(array)-1)

	fmt.Println(reflect.DeepEqual(array, valid))
}
