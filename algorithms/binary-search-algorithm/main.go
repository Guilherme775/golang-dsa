package main

import (
	"errors"
	"fmt"
	"math"
)

// Assuming that the array is ordered
func binarySearchAlgorithm(array []int, value int) (int, error) {
	lo := 0
	hi := len(array) - 1

	for lo <= hi {
		middleIndex := int(math.Floor(float64(lo+hi) / 2))
		middle := array[middleIndex]

		if middle == value {
			return middle, nil
		} else if middle > value {
			hi = middleIndex - 1
		} else {
			lo = middleIndex + 1
		}
	}

	return 0, errors.New("cannot find value")
}

func main() {
	slice := []int{2, 8, 16, 25, 35, 48, 92, 104, 205, 306, 411, 512}
	item, err := binarySearchAlgorithm(slice, 512)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(item)
}
