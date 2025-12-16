// todo: Преобразуйте переменную age и foo в число
// var := "23"
// foo := "23abc"
//
// Преобразуйте переменную age в Boolean
// age := "123abc"
//
// Преобразуйте переменную flag в Boolean
// flag := 1
//
// Преобразуйте значение в Boolean
// str_one := "Privet"
// str_two := ""
//
// # Преобразуйте значение 0 и 1 в Boolean
//
// Преобразуйте false в строку
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var age string = "1"
	var foo string = ""

	ageNum, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Ошибка")
	}

	ageBool, err := strconv.ParseBool(age)
	if err != nil {
		fmt.Println("Ошибка")
	}

	fmt.Printf("Nums %d, foo %s, bool, agebool %t", ageNum, foo, ageBool)
}
