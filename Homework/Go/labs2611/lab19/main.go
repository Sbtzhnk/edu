package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wGroup sync.WaitGroup
	wGroup.Add(2) // в группе две горутины
	doWork := func(id int) {
		defer wGroup.Done()
		fmt.Printf("Горутина %d начала выполнение \n", id)
		time.Sleep(3 * time.Second)
		fmt.Printf("Горутина %d завершила выполнение \n", id)
	}

	// вызываем горутины
	go doWork(1)
	go doWork(2)
	go doWork(3)

	wGroup.Wait() // ожидаем завершения обоих горутин
	fmt.Println("Горутины завершили выполнение")
}
