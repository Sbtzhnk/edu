// todo: Решить линейное уравнение A·x + B = 0, заданное своими коэффициентами A и B (коэффициент A не равен 0).
// Примечание: коэффициенты получаем через консольный ввод.
package main

import (
	"fmt"
)

func main() {
	for {
		var a, b int
		fmt.Println("Введите коэффициент A:")
		fmt.Scan(&a)

		if a == 0 {
			fmt.Println("Коэффициент A не должен быть равен нулю")
			continue
		}

		fmt.Println("Введите коэффициент B:")
		fmt.Scan(&b)

		fmt.Printf("X в уравнении %d*x + %d = 0 равен %d\n", a, b, -b/a)

		break
	}

}
