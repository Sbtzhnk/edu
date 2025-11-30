// В восточном календаре принят 60-летний цикл, состоящий из 12- летних подциклов,
// обозначаемых названиями цвета: зеленый, красный, желтый, белый и черный.
// В каждом подцикле годы носят названия животных: крысы, коровы, тигра, зайца, дракона,
// змеи, лошади, овцы, обезьяны, курицы, собаки и свиньи. По номеру года вывести его название,
// если 1984 год был началом цикла — годом зеленой крысы.
package main

import (
	"fmt"
)

func main() {
	var y int
	fmt.Println("Enter the year:")
	fmt.Scan(&y)
	start := 1984
	colors := []string{"green", "red", "yellow", "white", "black"}
	animals := []string{
		"rat", "cow", "tiger", "rabbit", "dragon", "snake",
		"horse", "goat", "monkey", "rooster", "dog", "pig",
	}
	diff := y - start

	colorIndex := (diff % 10) / 2
	animalIndex := diff % 12
	color := colors[colorIndex]
	animal := animals[animalIndex]

	fmt.Printf("This is %v %v year!\n", color, animal)

}
