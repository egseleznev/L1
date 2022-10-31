// Разработать программу, которая перемножает, делит, складывает, вычитает
// Две числовых переменных a,b, значение которых > 2^20

package main

import (
	"fmt"
	"math"
	"math/big"
)

func doCalculation(firstValue, secondValue *big.Float, operator string) *big.Float {

	result := new(big.Float)

	switch operator {
	// в зависимости от переданного оператора, вызываем соответствующую функцию

	case "+":
		return result.Add(firstValue, secondValue) // складываем значения с заданной точностью

	case "-":
		return result.Sub(firstValue, secondValue) // вычитаем одно значения из другого с заданной точностью

	case "*":
		return result.Mul(firstValue, secondValue) // перемножаем значения с заданной точностью

	case "/":
		return result.Quo(firstValue, secondValue) // делим одно значение на другое с заданной точностью

	default:
		return result // если оператор задан некорректно, возвращаем 0
	}
}
func main() {

	a, b := big.NewFloat(math.Pow(2, 21)), big.NewFloat(math.Pow(2, 25)) // инициализируем значения float >2^20, стандартная точность - 53, режим округления - до ближайшего целого

	fmt.Println(doCalculation(a, b, "+"))
	fmt.Println(doCalculation(a, b, "-"))
	fmt.Println(doCalculation(a, b, "*"))
	fmt.Println(doCalculation(a, b, "/"))

}
