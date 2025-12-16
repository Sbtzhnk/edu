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

func (n *Node) Next() *Node {
	return n.next
}

type Route struct {
	head   *Node
	length int
}

func (r *Route) insertAtHead(name string, time int, dist float32) {
	nodeHead := &Node{name, time, dist, nil}
	if r.head == nil {
		r.head = nodeHead
	} else {
		temp := r.head
		r.head = nodeHead
		nodeHead.next = temp
	}
	r.length += 1
}

func (r *Route) calcDistance(x *Node, y *Node) (int, float32) {
	if x == nil || y == nil || x == y {
		return 0, 0
	}

	var totalTime int
	var totalDist float32
	current := x

	// Проходим от узла X до узла, предшествующего Y
	for current != nil && current.next != nil && current != y {
		totalTime += current.time
		totalDist += current.dist
		current = current.next
	}

	return totalTime, totalDist
}

func (r *Route) display() {
	current := r.head
	for current != nil {
		fmt.Printf("Узел: %s, Время: %d, Расстояние: %.2f\n", current.name, current.time, current.dist)
		current = current.next
	}
}

func (r *Route) findNodeByName(name string) *Node {
	current := r.head
	for current != nil {
		if current.name == name {
			return current
		}
		current = current.next
	}
	return nil
}

func main() {
	list := Route{nil, 0}
	list.insertAtHead("C", 15, 18)
	list.insertAtHead("B", 5, 6)
	list.insertAtHead("A", 0, 0)

	fmt.Println("Весь маршрут:")
	list.display()
	fmt.Println()

	nodeA := list.findNodeByName("A")
	nodeB := list.findNodeByName("B")
	nodeC := list.findNodeByName("C")

	fmt.Println("calcDistance (исключая конечный узел):")
	time1, dist1 := list.calcDistance(nodeA, nodeC)
	fmt.Printf("A -> C: время=%d, расстояние=%.2f\n", time1, dist1)

	time2, dist2 := list.calcDistance(nodeB, nodeC)
	fmt.Printf("B -> C: время=%d, расстояние=%.2f\n", time2, dist2)

	time3, dist3 := list.calcDistance(nodeA, nodeB)
	fmt.Printf("A -> B: время=%d, расстояние=%.2f\n", time3, dist3)
}
