// todo: Единицы массы пронумерованы следующим образом: 1 — килограмм, 2 — миллиграмм, 3 — грамм,
//
//	4 — тонна, 5 — центнер. Дан номер единицы массы и масса тела M в этих единицах (вещественное число).
//	Вывести массу данного тела в килограммах
package main

import (
	"fmt"
)

type unitOfMeasure struct {
	id      uint8
	measure string
	factor  float32
}

func main() {
	var unit uint8
	var mass float32
	fmt.Println("Введите номер единицы измерения. 1 — килограмм, 2 — миллиграмм, 3 — грамм, 4 — тонна, 5 — центнер.")
	fmt.Scan(&unit)
	fmt.Println("Введите значение массы объекта.")
	fmt.Scan(&mass)

	measureList := []unitOfMeasure{
		{id: 1, measure: "Килограмм", factor: 1},
		{id: 2, measure: "Миллиграмм", factor: 0.000001},
		{id: 3, measure: "Грамм", factor: 0.001},
		{id: 4, measure: "Тонна", factor: 1000},
		{id: 5, measure: "Центнер", factor: 100},
	}

	for _, m := range measureList {
		if m.id == unit {
			fmt.Printf("Введено значение %v %v \nЭто %v килограмм.\n", mass, m.measure, mass*m.factor)
			break
		}
	}
}
