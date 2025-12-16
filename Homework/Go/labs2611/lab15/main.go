// todo: Дан целочисленный массив размера N из 10 элементов.
// Преобразовать массив, увеличить каждый его элемент на единицу.
package main

import (
	"fmt"
)

func main() {
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for idx := range array {
		newVal := array[idx] + 1
		array[idx] = newVal
		fmt.Println(newVal)
	}
}
