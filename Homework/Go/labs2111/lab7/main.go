// todo: Даны три точки A , B , C на числовой оси. Найти длины отрезков AC и BC и их сумму.
// Примечание: все точки получаем через консольный ввод.
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, ac, bc float64

	fmt.Println("Введите на каком расстояние от 0 находится точка А (на оси X):")
	fmt.Scan(&a)
	fmt.Println("Введите на каком расстояние от 0 находится точка B (на оси X):")
	fmt.Scan(&b)
	fmt.Println("Введите на каком расстояние от 0 находится точка C (на оси X):")
	fmt.Scan(&c)

	ac = math.Abs(c - a)
	bc = math.Abs(c - b)

	fmt.Printf("Длина отреза AC: %.5f,\nДлина отрезака BC: %.5f,\nСумма AC и BC: %.5f\n", ac, bc, ac+bc)

}
