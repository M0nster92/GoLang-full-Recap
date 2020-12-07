package main

import "fmt"

type List struct {
	head *Node
	tail *Node
}

type Node struct {
	value int
	next  *Node
	prev  *Node
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(value int) {
	node := &Node{value: value}
	if l.head == nil {
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
	}

	l.tail = node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

func (n *Node) Display() {
	for n != nil {
		fmt.Println(n.value)
		n = n.next
	}
}

func (n *Node) Reverse() {
	for n != nil {
		fmt.Println(n.value)
		n = n.prev
	}
}

func main() {
	l := &List{}
	l.Push(1)
	l.Push(2)
	l.Push(3)

	first := l.First()
	fmt.Println("First Node ", first.value)
	last := l.Last()
	fmt.Println("Last Node ", last.value)
	fmt.Println("Displaying Current Order: ")
	first.Display()
	fmt.Println("Displaying Reverse Order :")
	last.Reverse()

}
