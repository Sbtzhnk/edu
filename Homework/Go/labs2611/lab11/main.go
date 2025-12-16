// todo: Дан номер месяца (1 — январь, 2 — февраль, ...). Вывести название соответствующего
// времени года ("зима", "весна" и т.д.)
package main

import (
	"fmt"
)

type month struct {
	num    uint8
	name   string
	season string
}

func main() {
	var nm uint8
	fmt.Println("Введите номер месяца:")
	fmt.Scan(&nm)
	year := []month{
		{num: 1, name: "Январь", season: "Зима"},
		{num: 2, name: "Февраль", season: "Зима"},
		{num: 3, name: "Март", season: "Весна"},
		{num: 4, name: "Апрель", season: "Весна"},
		{num: 5, name: "Май", season: "Весна"},
		{num: 6, name: "Июнь", season: "Лето"},
		{num: 7, name: "Июль", season: "Лето"},
		{num: 8, name: "Август", season: "Лето"},
		{num: 9, name: "Сентябрь", season: "Осень"},
		{num: 10, name: "Осень", season: "Осень"},
		{num: 11, name: "Ноябрь", season: "Осень"},
		{num: 12, name: "Декабрь", season: "Зима"},
	}

	for _, m := range year {
		if m.num == nm {
			fmt.Printf("Месяц: %s\n", m.name)
			fmt.Printf("Время года: %s\n", m.season)
			break
		}
	}

}
