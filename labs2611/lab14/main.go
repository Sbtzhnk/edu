//todo: Дан массив размера N. Найти минимальное растояние между одинаковыми значениями в массиве и вывести их индексы.
// Одинаковых значение может быть два и более !
//Пример:
//var mass = [...]int{1, 2, 17, 54, 30, 89, 2, 1, 6, 2}

// Для числа 1 минимальное растояние в массиве по индексам: 0 и 7
// Для числа 2 минимальное растояние в массиве по индексам: 6 и 9
// Для числа 17 нет минимального растояния т.к элемент в массиве один.
package main

import (
	"fmt"
	"math"
)

func findMinDistance(arr []int) {
	valueIdxs := make(map[int][]int)

	for i, value := range arr {
		valueIdxs[value] = append(valueIdxs[value], i)
	}

	for value, idxs := range valueIdxs {
		if len(idxs) < 2 {
			fmt.Printf("Число %d одно.\n", value)
			continue
		}

		minDistance := math.MaxInt32
		var minIdxs [2]int

		for i := 0; i < len(idxs)-1; i++ {
			for j := i + 1; j < len(idxs); j++ {
				distance := idxs[j] - idxs[i]
				if distance < minDistance {
					minDistance = distance
					minIdxs[0] = idxs[i]
					minIdxs[1] = idxs[j]
				}
			}
		}

		fmt.Printf("Для числа %d минимальное расстояние в массиве по индексам: %d и %d\n",
			value, minIdxs[0], minIdxs[1])
	}
}

func main() {
	mass := []int{1, 2, 17, 54, 30, 89, 2, 1, 6, 2}

	fmt.Println("Массив:", mass)
	fmt.Println("Результаты:")
	findMinDistance(mass)
}
