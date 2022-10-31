// Разработать программу, которая будет последовательно отправлять значения в канал,
// А с другой стороны канала — читать
// По истечению N секунд программа должна завершаться
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	var workingTime time.Duration = 2 // время работы


	wg := new(sync.WaitGroup) // создаем вэйтгруппу для синхронизации горутин

	ch := make(chan interface{}) // инициализируем канал

	ctx, cancel := context.WithTimeout(context.Background(), workingTime*time.Second)	// создаем контекст, который завершится через заданное количество секунд
	defer cancel()

	wg.Add(1) // увеличиваем внутренний счетчик вэйтгруппы
	go func() {

	reading:
		for {

			select {

			case <-ctx.Done():	// если контекст завершился, закрываем канал и выходим
				close(ch)
				wg.Done()
				break reading

			case data := <-ch:
				fmt.Println(data)	// при получении данных из канала выводим их
			}

		}
	}()

	wg.Add(1)
	go func() {
		i := 0

	writing:
		for {

			i++

			select {

			case <-ctx.Done():	// если контекст завершился, закрываем канал и выходим
				wg.Done()
				break writing

			default:
				ch <- i	// записываем счетчик в канал
			}

		}
	}()

	wg.Wait() // ждем завершения горутин

}
