// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде
// По завершению программа должна выводить итоговое значение счетчика

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Counter представляет собой структуру-счетчик
type Counter struct {
	value uint64
}

// Атомарные операции - операции, завершающиеся в один шаг относительно других потоков, имеющих доступ к этой памяти
// Атомарные операции позволяют выполнить только одну операцию, в отличие, например, от лока и анлока мьютекса,
// Внутри которого можем выполнять несколько операций. Однако атомарные операции работают гораздо быстрее
func (c *Counter) IncValue(amount uint64) {
	atomic.AddUint64(&c.value, amount) // увеличиваем счетчик на amount
}

func (c *Counter) GetValue() uint64 {
	return atomic.LoadUint64(&c.value) // получаем значение счетчика
}

func main() {

	counter := Counter{}

	wg := new(sync.WaitGroup)

	for i := 0; i < 1000; i++ {

		wg.Add(1)

		go func() {

			for j := 0; j < 1000; j++ {
				counter.IncValue(1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println(counter.GetValue())
}
