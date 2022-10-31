// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации

package main

import (
	"strings"
)

var justString string

func createHugeString(length int) string {

	var sb strings.Builder

	for i := 0; i < length; i++ {
		sb.WriteString("б")
	}

	return sb.String()
}

func wrongSomeFunc() {

	v := createHugeString(1 << 10)
	justString = v[:100] // получим 100 байт, а не символов, так как срез идет по байтам, однако не все символы занимают 1 байт
}

func correctSomeFunc() {

	v := createHugeString(1 << 10)

	justString = string([]rune(v)[:100]) // получим 100 символов

}

func main() {
	correctSomeFunc()
}
