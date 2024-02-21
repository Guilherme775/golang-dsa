package main

import "fmt"

func filter(slice []int, conditionFn func(int) bool) []int {
	result := []int{}

	for i := 0; i < len(slice); i++ {
		if conditionFn(slice[i]) {
			result = append(result, slice[i])
		}
	}

	return result
}

func quickSort(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	if len(slice) == 1 {
		return slice
	}

	head := slice[0]
	tail := slice[1:]

	sm := filter(tail, func(i int) bool { return i < head })
	lg := filter(tail, func(i int) bool { return i > head })

	return append(append(quickSort(sm), head), quickSort(lg)...)
}

func main() {
	slice := []int{4, 77, 55, 2, 89, 21, 100, 23, 45, 78}

	fmt.Println(quickSort(slice))
}
