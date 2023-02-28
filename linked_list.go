package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) add(data int) {
	newNode := &Node{data, nil}
	fmt.Println("A", l.head)

	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		fmt.Println("B", current.next)
		for current.next != nil {

			current = current.next
			fmt.Println("B", current)
		}
		current.next = newNode

	}
}

func (l *LinkedList) print() {
	current := l.head
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Printf("\n")
}

func main() {
	list := LinkedList{}
	list.add(1)
	list.add(2)
	list.add(3)
	list.print()
}
