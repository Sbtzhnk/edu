package main

import (
	"fmt"
	"sync"
)

var counterGeneral int = 0 //  общий ресурс

func someWork(number int, channel chan bool, mutexCounter *sync.Mutex) {
	mutexCounter.Lock() // блокируем доступ к переменной counterGeneral
	counterGeneral = 0
	for k := 1; k <= 5; k++ {
		counterGeneral++
		fmt.Println("Goroutine", number, "-", counterGeneral)
	}
	mutexCounter.Unlock() // деблокируем доступ
	channel <- true
}

func main() {

	channel := make(chan bool)  // канал
	var mutexCounter sync.Mutex // определяем мьютекс
	for i := 1; i < 5; i++ {
		go someWork(i, channel, &mutexCounter)
	}

	for i := 1; i < 5; i++ {
		<-channel
	}

	fmt.Println("The End")
}
