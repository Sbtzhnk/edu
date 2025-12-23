package main

import (
	"fmt"
	"time"
)

func workRecreation(id int) {
	fmt.Println("", id, "snore ...")
	for j := id; j > 0; j-- {
		fmt.Printf("AAA")
		time.Sleep(2 * time.Second)
	}
	fmt.Print("BBB")
	time.Sleep(3 * time.Second)
}

func main() {
	for i := 0; i < 500; i++ {
		go workRecreation(i)
	}
	time.Sleep(5 * time.Second)
}
