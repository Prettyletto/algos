package main

import "fmt"

type Node[T any] struct {
	element T
	prev    *Node[T]
	next    *Node[T]
}

type LinkedList[T any] struct {
	lenght int
	head   *Node[T]
	tail   *Node[T]
}

func (l *LinkedList[T]) print() {
	l.printHelper(l.head)
}

func (l *LinkedList[T]) printHelper(curr *Node[T]) {
	if curr == nil {
		return
	}
	fmt.Println(curr.element)
	l.printHelper(curr.next)

}

func (l *LinkedList[T]) append(element T) {
	if l.head == nil {
		l.head = &Node[T]{element: element, prev: nil, next: nil}
		l.tail = l.head
		l.lenght++

		return
	}
	newNode := &Node[T]{element: element, prev: l.tail, next: nil}
	l.tail.next = newNode
	l.tail = newNode
	l.lenght++
}

func (l *LinkedList[T]) popHead() (T, error) {
	var blank T
	if l.lenght == 0 {
		return blank, fmt.Errorf("No values in the list, returning blank value")
	}

	tmp := l.head
	if l.head.next == nil {
		l.head = nil
		l.tail = nil
		l.lenght--
		return tmp.element, nil
	}
	l.head = l.head.next
	l.lenght--
	return tmp.element, nil
}

func (l *LinkedList[T]) popTail() (T, error) {
	var blank T
	if l.lenght == 0 {
		return blank, fmt.Errorf("No values in the list, returning blank value")
	}

	tmp := l.tail
	if l.tail.prev == nil {
		l.tail = nil
		l.head = nil
		l.lenght--
		return tmp.element, nil
	}
	l.tail = l.tail.prev
	l.tail.next = nil
	l.lenght--

	return tmp.element, nil
}

func (l *LinkedList[T]) prepend(element T) {
	if l.head == nil {
		l.head = &Node[T]{element: element, prev: nil, next: nil}
		l.tail = l.head
		l.lenght++

		return
	}
	newNode := &Node[T]{element: element, prev: nil, next: l.head}
	l.head.prev = newNode
	l.head = newNode
	l.lenght++

}

func (l *LinkedList[T]) getAt(index int) (T, error) {
	var blank T
	if index < 0 || index >= l.lenght {
		return blank, fmt.Errorf("Index out of boundaries")
	}
	curr := l.head
	for i := 0; i <= index; i++ {
		if i == index {
			return curr.element, nil
		}
		curr = curr.next

	}

	return blank, nil
}

func (l *LinkedList[T]) insertAt(element T, index int) {
	if index < 0 || index > l.lenght {
		panic("index out of bounds")
	}

	if index == l.lenght {
		l.append(element)
		return
	}

	if index == 0 {
		l.prepend(element)
		return
	}
	curr := l.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}
	newNode := &Node[T]{element: element, prev: curr.prev, next: curr}
	if curr.prev != nil {
		curr.prev.next = newNode
	} else {
		l.head = newNode
	}
	curr.prev = newNode
	l.lenght++

}
