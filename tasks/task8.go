// Дана переменная int64
// Разработать программу которая устанавливает i-й бит в 1 или 0

package main

import "fmt"

func changeBit(value int64, position int, toZero bool) int64 {

	if !toZero {
		return value | int64(1<<position) // двигаем бит в маске и с помощью операции логического или меняем бит в переменной
	} else {
		return value &^ int64(1<<position) // двигаем бит в маске, инвертируем маску и с помощью операции логического и меняем бит в переменной
	}

}

func main() {

	var value int64 = 5624562436

	fmt.Println(changeBit(value, 3, false) == 5624562444)
	fmt.Println(changeBit(value, 14, false) == 5624578820)
	fmt.Println(changeBit(value, 8, true) == 5624562180)
	fmt.Println(changeBit(value, 30, true) == 4550820612)
}
