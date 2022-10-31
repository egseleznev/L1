// Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
// C инкапсулированными параметрами x,y и конструктором

package main

import (
	"fmt"
	"math"
)

// Структура представляет собой точку в двумерной системе координат
type Point struct {
	x float64
	y float64
}

// Конструктор, возвращающий указатель на объект точки
func newPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Возвращаем расчитанное по формуле расстояние между точками
func (p *Point) calculateDistance(pointSecond *Point) float64 {
	return math.Sqrt(math.Pow(pointSecond.x-p.x, 2) + math.Pow(pointSecond.y-p.y, 2))
}

func main() {
	pointFirst := newPoint(10.5, 12.25)
	pointSecond := newPoint(2, 5.75)

	fmt.Println(pointFirst.calculateDistance(pointSecond) == pointSecond.calculateDistance(pointFirst))
}
