package main

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head *Node[T]
}

func New[T comparable]() *LinkedList[T] {
	list := &LinkedList[T]{}

	return list
}

func (list *LinkedList[T]) Display() {
	currentNode := list.head

	if currentNode == nil {
		fmt.Println("Nil")
		return
	}

	for currentNode != nil {
		fmt.Println(fmt.Sprintf("%v", currentNode.value))
		currentNode = currentNode.next
	}
}

func (list *LinkedList[T]) Append(value T) {
	currentNode := list.head

	if currentNode == nil {
		list.head = &Node[T]{value: value}
		return
	}

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = &Node[T]{value: value}
}

func (list *LinkedList[T]) Prepend(value T) {
	currentHead := list.head
	list.head = &Node[T]{value: value, next: currentHead}
}

func (list *LinkedList[T]) Get(value T) (result T, err error) {
	currentNode := list.head

	for currentNode != nil {
		if currentNode.value == value {
			return currentNode.value, nil
		}

		currentNode = currentNode.next
	}

	var zero T
	return zero, errors.New("Cannot find value")
}

func (list *LinkedList[T]) Remove(value T) {
	if list.head != nil && list.head.value == value {
		list.head = list.head.next
	}

	currentNode := list.head
	var previousNode *Node[T]

	for currentNode != nil {
		if currentNode.value == value {
			if currentNode.next != nil {
				currentNode.value = currentNode.next.value
				currentNode.next = currentNode.next.next
			} else {
				previousNode.next = nil
			}

			break
		}

		previousNode = currentNode
		currentNode = currentNode.next
	}
}

func main() {
	list := New[string]()
	list.Append("append test")
	list.Append("append test2")
	list.Append("append test3")
	list.Prepend("prepend test")
	list.Prepend("prepend test2")
	list.Remove("append test3")
	list.Remove("prepend test2")
	result, err := list.Get("append test")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(result)
	fmt.Println("-----------")
	list.Display()
}
