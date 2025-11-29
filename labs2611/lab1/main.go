// todo: Дан номер месяца (1 — январь, 2 — февраль, ...). Вывести название соответствующего
// времени года ("зима", "весна" и т.д.)
package main

import (
	"fmt"
)

func GetMonth(ord uint8) string {
	if ord == 1 {
		return "Январь"
	}
	if ord == 2 {
		return "Февраль"
	}
	if ord == 3 {
		return "Март"
	}
	if ord == 4 {
		return "Апрель"
	}
	if ord == 5 {
		return "Май"
	}
	if ord == 6 {
		return "Июнь"
	}
	if ord == 7 {
		return "Июль"
	}
	if ord == 8 {
		return "Август"
	}
	if ord == 9 {
		return "Сентябрь"
	}
	if ord == 10 {
		return "Октябрь"
	}
	if ord == 11 {
		return "Ноябрь"
	}
	if ord == 12 {
		return "Декабрь"
	}
	return "Неизвестен"
}

func main() {
	var ord uint8
	var month string

	fmt.Println("Введите номер месяца:")
	fmt.Scan(&ord)

	if (ord < 1) || (ord > 12) {
		fmt.Println("Значение некорректно")
		return
	}

	month = GetMonth(ord)
	fmt.Printf("Название введённого месяца: %s\n", month)
}
