// Реализовать собственную функцию sleep

package main

import (
	"fmt"
	"time"
)

func sleepFirst(seconds int) {

	timeStamp := time.Now().UnixNano()

	for {
		if timeStamp+int64(time.Duration(seconds)*time.Second) <= time.Now().UnixNano() { // крутимся в цикле, пока указанное время не выйдет
			break
		}
	}
}

func sleepSecond(seconds int) {

	t := time.NewTimer(time.Duration(seconds) * time.Second) // создаем таймер, который вернет текущее время через указанное количество времени
	<-t.C                                                    // получаем текущее время, значит указанное время прошло
	t.Stop()                                                 // завершаем таймер
}

func main() {

	sleepFirst(2)

	fmt.Println("2 sec have passed")

	sleepSecond(1)
	fmt.Println("1 sec have passed")

}
