package main

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Stack[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func New[T any](value T) Stack[T] {
	return Stack[T]{head: &Node[T]{value: value}, length: 1}
}

func (stack *Stack[T]) Display() {
	currentNode := stack.head
	condition := true

	for condition {
		if currentNode != nil {
			fmt.Println(currentNode.value)
			currentNode = currentNode.next
		} else {
			condition = false
		}
	}
}

func (stack *Stack[T]) Push(value T) {
	newNode := &Node[T]{value: value}

	if stack.tail == nil {
		stack.head.next = newNode
		stack.tail = newNode
	} else {
		stack.tail.next = newNode
		stack.tail = newNode
	}

	stack.length += 1
}

func (stack *Stack[T]) Pop() (T, error) {
	var zero T

	if stack.head == nil {
		return zero, errors.New("empty stack")
	}

	node := stack.head

	stack.head = stack.head.next

	stack.length -= 1
	return node.value, nil
}

func (stack *Stack[T]) Peek() (T, error) {
	var zero T

	if stack.head == nil {
		return zero, errors.New("empty stack")
	}

	return stack.head.value, nil
}

func main() {
	stack := New(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Pop()
	stack.Display()
}
