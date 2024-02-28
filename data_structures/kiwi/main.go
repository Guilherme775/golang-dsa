package main

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type Kiwi[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func New[T any](value T) Kiwi[T] {
	return Kiwi[T]{head: &Node[T]{value: value}, length: 1}
}

func (kiwi *Kiwi[T]) Display() {
	currentNode := kiwi.head
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

func (kiwi *Kiwi[T]) Enqueue(value T) {
	newNode := &Node[T]{value: value}

	if kiwi.tail == nil {
		kiwi.head.next = newNode
		kiwi.tail = newNode
	} else {
		kiwi.tail.next = newNode
		kiwi.tail = newNode
	}

	kiwi.length += 1
}

func (kiwi *Kiwi[T]) Dequeue() (T, error) {
	var zero T

	if kiwi.head == nil {
		return zero, errors.New("empty queue")
	}

	head := kiwi.head

	kiwi.head = kiwi.head.next

	kiwi.length -= 1

	return head.value, nil
}

func (kiwi *Kiwi[T]) Peek() (T, error) {
	var zero T

	if kiwi.head == nil {
		return zero, errors.New("empty queue")
	}

	head := kiwi.head

	return head.value, nil
}

func main() {
	kiwi := New(1)
	kiwi.Enqueue(2)
	kiwi.Enqueue(3)
	kiwi.Enqueue(4)
	kiwi.Dequeue()
	kiwi.Display()
}
