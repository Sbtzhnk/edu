// todo: Определить в коде переменные:
// 1. Целочисленного типа
// 2. Вещественного типа
// 3. Логического типа
// 4. Строкового типа
// 5. Пустого типа
// 6. Задайте указатели на все перечисленные типы.
// Вывести их типы (надо погуглить).

package main

import (
	"fmt"
)

func main() {
	var amount int = 0
	var price float32 = 0.32
	var isAllowed bool = true
	var code string = "111"

	var amountPtr *int = &amount
	var pricePtr *float32 = &price
	var isAllowedPtr *bool = &isAllowed
	var codePtr *string = &code

	fmt.Printf("Тип %T, Значение %d, Указатель %p\n", amount, amount, amountPtr)
	fmt.Printf("Тип %T, Значение %.5f, Указатель %p\n", price, price, pricePtr)
	fmt.Printf("Тип %T, Значение %t, Указатель %p\n", isAllowed, isAllowed, isAllowedPtr)
	fmt.Printf("Тип %T, Значение %v, Указатель %p\n", code, code, codePtr)

}
