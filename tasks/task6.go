// Реализовать все возможные способы остановки выполнения горутины

package main

import (
	"context"
	"fmt"
	"sync/atomic"
)

var goroutinesAmount int32 // счетчик горутин в работе для проверки корректности

func closingChannel() {

	ch := make(chan bool)

	go func() {

		atomic.AddInt32(&goroutinesAmount, 1)
		defer atomic.AddInt32(&goroutinesAmount, -1)

		for {
			select {

			case <-ch:
				return // горутина завершит работу при получении сигнала из канала

			default:

			}
		}
	}()

	ch <- true
}

func rangeChannel() {

	ch := make(chan struct{}, 1)

	go func() {

		atomic.AddInt32(&goroutinesAmount, 1)
		defer atomic.AddInt32(&goroutinesAmount, -1)

		for range ch {
		}
		// при закрытии канала, горутина завершит работу

	}()

	close(ch)
}

func cancelContext() {

	ctx, cancel := context.WithTimeout(context.Background(), 100) // контекст с таймаутом

	go func() {

		atomic.AddInt32(&goroutinesAmount, 1)
		defer atomic.AddInt32(&goroutinesAmount, -1)

		for {
			select {

			case <-ctx.Done(): // при завершении контекста, горутина завершит работу
				return

			default:
			}
		}
	}()

	cancel()
}

func withPanic() {
	go func() {
		atomic.AddInt32(&goroutinesAmount, 1)

		defer func() {
			atomic.AddInt32(&goroutinesAmount, -1)
			recover()
		}()

		panic("expected error") // при возникновении паники горутина прервет свою работу, выполнится recover если вызван
	}()
}

func main() {

	closingChannel()
	rangeChannel()
	cancelContext()
	withPanic()

	fmt.Println(goroutinesAmount)

}
