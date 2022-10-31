// Реализовать постоянную запись данных в канал (главный поток)
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout
// Необходима возможность выбора количества воркеров при старте
// Программа должна завершаться по нажатию Ctrl+C
// Выбрать и обосновать способ завершения работы всех воркеров

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
)

// запускаем N воркеров, которые читают данные из канала, пока он не закроется, таким образом обеспечивается graceful shutdown
func runWorkers(wg *sync.WaitGroup, workersAmount int, ch chan int) {

	for i := 0; i < workersAmount; i++ {

		wg.Add(1)
		go func() {

			for str := range ch {
				fmt.Println(str)
			}
			wg.Done()
		}()
	}
}

func main() {

	var workersAmount int // переменная для хранения количества воркеров

	fmt.Print("Enter amount of workers: ")

	if _, err := fmt.Scanf("%d", &workersAmount); err != nil { // ждем ввода количества воркеров пользователем
		log.Fatal(err)
	}

	var wg sync.WaitGroup // создаем вэйтгруппу для синхронизации горутин

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt) // создаем контекст, который отслеживает прерывание
	defer cancel()

	ch := make(chan int) // создаем канал для записи данных

	wg.Add(1) // увеличиваем внутренний счетчик вэйтгруппы
	go func() {
		i := 0

	writing:
		for {

			select { // ожидаем операции

			case <-ctx.Done(): // при получении прерывания контекст завершается, канал закрыватся, так как канал закрылся, воркеры тоже завершат работу
				close(ch)
				break writing

			default: // пока контекст жив, пишем в канал значение счетчика
				ch <- i
			}

			i++
		}

		wg.Done() // говорим что горутина завершила работу
	}()

	runWorkers(&wg, workersAmount, ch) // запускаем воркеров, которые читают данные из канала
	wg.Wait()                          // ждем завершения горутин
}
