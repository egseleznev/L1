// Реализовать конкурентную запись данных в map

package main

import (
	"fmt"
	"sync"
)

// Структура конкуретной мапы
type syncMap struct {
	mx        sync.Mutex
	unsyncMap map[int]interface{}
}

// Конструктор
func (s *syncMap) create() *syncMap {

	s.unsyncMap = make(map[int]interface{})

	s.mx = sync.Mutex{}

	return s
}

func (s *syncMap) getValue(key int) (interface{}, bool) {

	s.mx.Lock() // лочим мьютекс перед тем как прочитать значение

	v, ok := s.unsyncMap[key] // читаем значение

	s.mx.Unlock() // после прочтения разблокируем ресурс

	return v, ok
}

func (s *syncMap) loadValue(key int, value interface{}) {

	s.mx.Lock() // лочим мьютекс перед тем как записать значение

	s.unsyncMap[key] = value // пишем значение

	s.mx.Unlock() // после записи разблокируем ресурс
}

func main() {

	newMap := (&syncMap{}).create()

	go func() {
		newMap.loadValue(1, "str")
	}()

	go func() {
		newMap.loadValue(2, 5)
	}()

	go func() {
		newMap.loadValue(3, false)
	}()

	fmt.Println(newMap.getValue(1))
	fmt.Println(newMap.getValue(2))
	fmt.Println(newMap.getValue(3))
}
