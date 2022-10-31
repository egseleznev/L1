// Реализовать паттерн «адаптер» на любом примере

package main

import (
	"fmt"
	"strconv"
)

// Предположим, есть сервис, работающий только со строковыми данными
type sendStringData interface {
	sendString()
}

// тип строковых данных
type stringData struct {
	data string
}

// имплементация интерфейса, отправка строковых данных
func (s *stringData) sendString() {
	fmt.Printf("%T %v \n", s.data, s.data)
}

// далее у нас появилась необходимость работы с целочисленными данными
type integerData struct {
	data int
}

// однако, вышеупомянутый сервис работает только со строками и, дабы не писать повторяющийся код, воспользуемся
// паттерном "адаптер", для этого реализуем функцию конвертации целого числа к строке, а также структуру адаптера

func (i *integerData) convertToString() string {
	return strconv.Itoa(i.data)
}

type integerDataAdapter struct {
	integerData *integerData
}

// с помощью адаптера появилась возможность отправлять целочисленные данные в строковом формате, реализовав через него интерфейс
func (adapter *integerDataAdapter) sendString() {

	dataToSend := adapter.integerData.convertToString()
	fmt.Printf("%T %v \n", dataToSend, dataToSend)
}

func main() {

	strData := &stringData{data: "5"}
	strData.sendString()

	intData := &integerData{data: 5}

	adapter := &integerDataAdapter{intData}
	adapter.sendString()
}
