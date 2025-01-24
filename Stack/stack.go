package main

type Node[T any] struct {
	element T
	prev    *Node[T]
}

type Stack[T any] struct {
	length int
	head   *Node[T]
}

func (s *Stack[T]) push(element T) {
	newNode := &Node[T]{element: element, prev: nil}
	s.length++
	if s.length == 0 {
		s.head = newNode
		return
	}
	newNode.prev = s.head
	s.head = newNode

}

func (s *Stack[T]) pop() T {
	var BLANK T
	if s.length == 0 {
		return BLANK
	}

	h := s.head
	s.head = s.head.prev
	return h.element

}

func (s *Stack[T]) peek() T {
	var BLANK T
	if s.length == 0 {
		return BLANK
	}

	return s.head.element
}
