package main

import (
	"fmt"
	"time"
)

func workRecreation(id int, channel chan int) { // канал - аргумент
	fmt.Println("Recreation ...", id, "_workerId")
	time.Sleep(4 * time.Second) // работник отдыхает
	channel <- id               // отправим значение через канал в main
}

func main() {
	channel := make(chan int) // создаем канал для связи с типом объекта int
	for i := 0; i < 11; i++ {
		go workRecreation(i, channel)
	}
	for i := 0; i < 11; i++ {
		workerId := <-channel // получаем значение и присваиваем его переменной 'workerId'
		fmt.Println("worker ", workerId, " finished resting")
	}
}
