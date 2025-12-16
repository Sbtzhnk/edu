// todo: Данные две переменные:
// age := 36.6
// temperature := 25
// Нужно обменять значения переменных местами. В итого age
// должен равнятся 25 а temperature – 36.6
package main

import (
	"fmt"
)

func main() {
	fmt.Print("ok")
	var age float32 = 36.6
	var temperature float32 = 25
	age, temperature = temperature, age
	fmt.Println(age, temperature)
}
