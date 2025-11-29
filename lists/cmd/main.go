package main

import "fmt"

type Stack struct {
	list []int
}

func (s *Stack) Pop() int {
	i := len(s.list)
	elm := s.list[i-1]
	s.list = s.list[:i-1]
	return elm
}

func (s *Stack) Size() int {
	return len(s.list)
}

func (s *Stack) Push(val int) {
	s.list = append(s.list, val)
}

func main() {
	mass := []int{1, 2, 3}
	st := Stack{mass}
	fmt.Println(st.Pop()):~/edu/lab2911/singlton
	fmt.Println(mass)
	fmt.Println(st.Pop())
	st.Push(5)
	fmt.Println(mass)
	fmt.Println(st.list)
}
