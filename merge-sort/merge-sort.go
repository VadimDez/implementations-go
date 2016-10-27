package main;

import "fmt"

func main() {
	var a = []int{9, 8, 3, 6}


	fmt.Println("Input:")
	fmt.Println(a)
	fmt.Println("Result")
	fmt.Println(mergeSort(a, len(a)))
}

func mergeSort(a []int, n int) []int {
	if n <= 1 {
		return a
	}

	var middle = n/2;
	return merge(mergeSort(a[0:middle], middle), mergeSort(a[middle:], n - middle))
}

func merge(a []int, b []int) []int {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	if a[0] < b[0] {
		return append([]int{a[0:1][0]}, merge(a[1:], b)...)
	}
	return append([]int{b[0:1][0]}, merge(a, b[1:])...)
}