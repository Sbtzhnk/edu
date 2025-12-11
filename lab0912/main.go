package main

import (
	"fmt"
	"time"
)

/*func zero(int ,int) {
	fmt.Println("zero")
}

func zero_(x int, y int) {
	fmt.Println("zero")
}

func zero__(x, y int) int {
	fmt.Println("zero")
	return x + y
}

func _zero___(x, y int) (int, string) {
	fmt.Println(zero)
	return x + y, "str"
}*/

func zero___(x, y int) (z int) {
	fmt.Println("zero")
	z = x + y
	return x
}

func printFunction(fn func(int, int) int) {
	x := 100
	y := 200
	summ := fn(x, y)

	fmt.Println(summ)
}

func newAnon(str string) func(int, int) int {
	fn := func(x int, y int) int {
		fmt.Println(str)
		return x + y
	}
	return fn
}

func Hello() string {
	fmt.Println("Hello function")
	return "Hello"
}

type tpFnSumm func(int, int) int

func DecoratorFN(fn tpFnSumm) func(int, int) {
	return func(x int, y int) {
		start := time.Now()
		res := fn(x, y)
		fmt.Println(res)
		elapsed := time.Since(start)
		fmt.Println(elapsed)
	}
}

func summ(x int, y int) int {
	return x + y
}

func main() {
	defer func() { fmt.Println(Hello()) }()
	defer fmt.Println("Stage 1")
	//zero(2, 5)
	//fmt.Println(zero___(2, 5))
	fn := func(x int, y int) int {
		fmt.Println("call")
		return x + y
	}
	fn(2, 4)
	newAnon("111")

	printFunction(fn)
	fn2 := DecoratorFN(summ)
	fn2(10, 10)
	panic("Fatal Error")
}
