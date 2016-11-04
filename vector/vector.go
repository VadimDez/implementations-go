package main

import "fmt"

type Vector struct {
	head *[]int
	numberOfItems int
	arrayCapacity int
}

func NewVector() *Vector {
	var a []int;
	return &Vector{&a, 0, 8};
}

func (v Vector) size() int {
	return v.numberOfItems;
}

func (v Vector) capacity() int {
	return v.arrayCapacity;
}

func main() {
	v := NewVector();
	fmt.Println("Size of v", v.size());

	fmt.Println("Capacity of v", v.capacity());
}
