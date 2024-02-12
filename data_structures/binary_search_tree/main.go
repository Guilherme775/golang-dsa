package main

import (
	"errors"
	"fmt"
)

type BinarySearchTree struct {
	value int
	left  *BinarySearchTree
	right *BinarySearchTree
}

func New(value int) *BinarySearchTree {
	tree := &BinarySearchTree{}
	tree.value = value

	return tree
}

func (tree *BinarySearchTree) Display(suffix string) {
	fmt.Println(suffix, tree.value)
	if tree.right != nil {
		tree.right.Display("right")
	}
	if tree.left != nil {
		tree.left.Display("left")
	}
}

func (tree *BinarySearchTree) Insert(value int) {
	currentNode := tree

	for currentNode != nil {
		if value > currentNode.value {
			if currentNode.right == nil {
				currentNode.right = &BinarySearchTree{value: value}
				break
			} else {
				currentNode = currentNode.right
			}
		} else if value < currentNode.value {
			if currentNode.left == nil {
				currentNode.left = &BinarySearchTree{value: value}
				break
			} else {
				currentNode = currentNode.left
			}
		}
	}
}

func (tree *BinarySearchTree) Search(value int) (result int, err error) {
	currentNode := tree

	for currentNode != nil {
		if currentNode.value == value {
			return currentNode.value, nil
		}

		if value > currentNode.value {
			currentNode = currentNode.right
		} else if value < currentNode.value {
			currentNode = currentNode.left
		}
	}

	return 0, errors.New("Cannot find value")
}

func main() {
	tree := New(0)
	tree.Insert(1)
	tree.Insert(2)
	tree.Insert(-1)
	tree.Insert(-2)
	result, err := tree.Search(0)

	if err != nil {
		fmt.Printf("Cannot find value")
		return
	}

	fmt.Println(result)
	fmt.Println("------------")
	tree.Display("root")
}
