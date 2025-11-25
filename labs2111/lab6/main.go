// todo: Дана сторона квадрата a. Найти его площадь S = a²
// Примечание: сторону квадрата получаем через консольный ввод.
package main

import (
	"fmt"
)

func main() {
	var a uint32

	fmt.Println("Введите сторону квадрата:")
	fmt.Scan(&a)
	fmt.Printf("Площадь: %d\n", a*a)
}
