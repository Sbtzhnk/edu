
// todo: Реализовать метод Next() для структуры Node !!!
// Реализовать метод подсчёта времени и расстояния  calcDistance( x Node, y Node) от узла X до узла Y для структуры Route.
/*
package main

import (
	"fmt"
)

type Node struct {
    name string
	time int
    dist float32
	next *Node
}

type Route struct {
    head *Node
    length int
}


func (r *Route) insertAtHead(name string, time int, dist float32 ) {
    nodeHead := &Node{name, time , dist  , nil}
    if r.head == nil {
        r.head = nodeHead
    } else {
        temp := r.head
        r.head = nodeHead
        temp.next = temp
    }
    r.length += 1
}


func main() {
    list := Route{nil, 0}
	list.insertAtHead("С", 15, 18)
	list.insertAtHead("B", 5,  6)
	list.insertAtHead("A", 0,  0)
	fmt.Println(list.head)
}
*/