package main

import (
	"fmt"
)

func main() {
	var a []int
	var b []int = []int{1, 2, 3}
	c := []int{1, 2, 3}
	d := []int{1: 12}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
