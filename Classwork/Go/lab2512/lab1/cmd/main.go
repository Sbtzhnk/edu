package main

import (
	"fmt"
	"lab2512lab1/pkg/bubble"
	"lab2512lab1/pkg/insert"
	"math/rand"
)

func main() {
	arr := []int{}
	for i := 0; i < 10; i++ {
		randomNum := rand.Intn(100)
		arr = append(arr, randomNum)
	}

	fmt.Println("До сортировки:", arr)
	bubble.BubbleSort(arr)
	fmt.Println("После сортировки пузырьком:", arr)
	arr = []int{64, 34, 25, 12, 22, 11, 90}
	insert.InsertionSort(arr)
	fmt.Println("После сортировки вставками:", arr)
}
