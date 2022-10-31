// Реализовать пересечение двух неупорядоченных множеств

package main

import "fmt"

// Создаем множество целочисленных значений
// Почти то же самое, что и хеш-таблица, но без сохранения значения
// Используем если нам нужна только быстрая проверка вхождения
// В качестве значения нужно лишь указать пустое значение struct{}, занимающее 0 байт
func createSet(sequence []int) map[int]struct{} {

	set := make(map[int]struct{})

	for _, v := range sequence {
		set[v] = struct{}{}
	}
	return set
}

func intersect(firstSet, secondSet map[int]struct{}) map[int]struct{} {

	result := createSet([]int{})

	for i, _ := range firstSet { // проходимся по значениям первого множество
		if _, ok := secondSet[i]; ok == true { // и ищем это значение во втором множестве
			result[i] = struct{}{} // если нашли, значит элемент входит в пересечение
		}
	}

	return result

}

func main() {

	firstSet := createSet([]int{1, 2, 3, 4, 5})
	secondSet := createSet([]int{2, 6, 7, 3, 6})

	fmt.Println(intersect(firstSet, secondSet))

}
