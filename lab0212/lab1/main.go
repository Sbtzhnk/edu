package main

import "fmt"

func (p *Person) getName() string {
	return p.name
}

func (p *Account) getName() string {
	return p.name
}

type intPerson interface {
	getName() string
}

type Person struct {
	name string
}

type Account struct {
	name string
}

func showVar(x intPerson) {
	fmt.Println(x)
	x.getName()
	switch x.(type) {
	case *Person:
		fmt.Println("Person")
	case *Account:
		fmt.Println("Account")
	default:
		fmt.Println("Empty")
	}
}

func main() {
	var x intPerson
	var y Person = Person{"s"}
	var ac Account
	x = &y
	showVar(&y)
	showVar(&ac)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
