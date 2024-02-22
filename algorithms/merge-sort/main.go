package main

import "fmt"

func splitAt[T any](s []T, index int) ([]T, []T) {
	if index < 0 || index > (len(s)-1) {
		index = len(s) - 1
	}

	return s[:index], s[index:]
}

func merge(xs []int, ys []int) []int {
	if len(xs) >= 1 && len(ys) == 0 {
		return xs
	}

	if len(ys) >= 1 && len(xs) == 0 {
		return ys
	}

	xsHead := xs[0]
	xsTail := xs[1:]
	ysHead := ys[0]
	ysTail := ys[1:]

	if xsHead < ysHead {
		return append([]int{xsHead}, merge(xsTail, ys)...)
	} else {
		return append([]int{ysHead}, merge(xs, ysTail)...)
	}
}

func mergeSort(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	if len(slice) == 1 {
		return slice
	}

	midpoint := len(slice) / 2
	xs, ys := splitAt(slice, midpoint)

	return merge(mergeSort(xs), mergeSort(ys))
}

func main() {
	slice := []int{4, 77, 55, 2, 89, 21, 100, 23, 45, 78}

	fmt.Println(mergeSort(slice))
}
