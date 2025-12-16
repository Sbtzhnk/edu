package main

// todo: Проверить истинность высказывания:
//"Данное четырехзначное число читается одинаково слева направо и справа налево".
import (
	"fmt"
	"strconv"
)

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	for {
		var a int
		var oa string
		fmt.Println("Введите четырёхзначное число:")
		fmt.Scan(&a)

		if (a < 1000) || (a > 9999) {
			fmt.Println("Число выходит за рамки задачи")
			continue
		}

		oa = strconv.Itoa(a)
		if oa == Reverse(oa) {
			fmt.Printf("%s число читается одинаково слева направо и справа налево \n", oa)
		} else {
			fmt.Printf("%s число не читается одинаково слева направо и справа налево \n", oa)
		}
		break
	}
}
