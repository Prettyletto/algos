package main

type Node[T any] struct {
	element T
	next    *Node[T]
}

type Queue[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func (q *Queue[T]) enqueue(element T) {
	q.length++
	newNode := &Node[T]{element: element, next: nil}
	if q.length == 0 {
		q.head = newNode
		q.tail = q.head
		return
	}
	q.tail.next = newNode
	q.tail = newNode

}

func (q *Queue[T]) dequeue() T {
	var BLANK T
	if q.length == 0 {
		return BLANK
	}
	q.length--
	h := q.head
	q.head = q.head.next

	return h.element
}

func (q *Queue[T]) peek() T {
	var BLANK T
	if q.head == nil {
		return BLANK
	}

	return q.head.element
}
