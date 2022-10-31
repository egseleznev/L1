// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}

package main

import "fmt"

// пустой интерфейс может принимать значение переменной любого типа
func determineType(value interface{}) string {

	switch value.(type) {
	// используем встроенное средство go для определения типа переменной

	case int:
		return fmt.Sprintf("%T \n", value)

	case string:
		return fmt.Sprintf("%T \n", value)

	case bool:
		return fmt.Sprintf("%T \n", value)

	case chan string:
		return fmt.Sprintf("%T \n", value)

	default:
		return fmt.Sprintf("unforeseen type: %T \n", value)
	}
}

func main() {

	intValue := 1
	strValue := "value"
	boolValue := true
	chanIntValue := make(chan string)
	floatValue := 20.5

	fmt.Println(determineType(intValue), determineType(strValue), determineType(boolValue), determineType(chanIntValue), determineType(floatValue))

}
