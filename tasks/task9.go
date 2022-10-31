// Разработать конвейер чисел
// Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2,
// После чего данные из второго канала должны выводиться в stdout

package main

import (
	"fmt"
	"sync"
)

func main() {

	array := []int{1, 2, 3, 4, 5} // массив чисел

	chNumbers, chSquares := make(chan int), make(chan int) // инициализируем два небуферизированных канала

	wg := new(sync.WaitGroup) // создаем вэйтгруппу для синхронизации горутин

	wg.Add(1) // увеличиваем внутренний счетчик вэйтгруппы
	go func(chNumbers chan int, array []int) {

		for _, v := range array {
			chNumbers <- v // пишем в канал 1 значение x
		}

		close(chNumbers) // когда все числа записали, закрываем канал

		wg.Done() // говорим что горутина завершила работу

	}(chNumbers, array)

	wg.Add(1)
	go func(chNumbers, chSquares chan int) {

		for v := range chNumbers {
			chSquares <- v * v // получаем значение x из канала 1 и пишем значение x^2 в канал 2
		}

		close(chSquares) // когда канал 1 закрылся, закрываем и 2, так как значения больше не придут

		wg.Done()

	}(chNumbers, chSquares)

	wg.Add(1)
	go func(chSquares chan int) {

		for data := range chSquares { // получаем значения x^2 из канала 2 и выводим их, пока канал не будет закрыт
			fmt.Println(data)
		}

		wg.Done()
	}(chSquares)

	wg.Wait() // ждем завершения всех горутин
}
