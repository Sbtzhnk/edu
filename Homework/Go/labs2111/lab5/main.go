// todo: Написать программу, которая считывает два числа и выводит их сумму, разность, частное, произведение,
// результат целочисленного деления, результат деления с остатком, результат возведения в степень.
package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(math.Pow(float64(a), float64(b)))
	fmt.Println(a % b)

}
