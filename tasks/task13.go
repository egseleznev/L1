// Поменять местами два числа без создания временной переменной
package main

import (
	"fmt"
)

func main() {
	firstValue, secondValue := 1, 2

	fmt.Println(firstValue, secondValue)

	// представлен математический способ замены местами двух чисел, также в go имеется встроенный вариант с синтаксисом a, b = b, a
	firstValue += secondValue

	secondValue = firstValue - secondValue

	firstValue = firstValue - secondValue

	fmt.Println(firstValue, secondValue)
}
