package main

import (
	"fmt"
	"errors"
)

type Vector struct {
	head []int
	numberOfItems int
	arrayCapacity int
}

func NewVector() *Vector {
	var a = []int{};
	var v = Vector{};

	v.numberOfItems = 0;
	v.arrayCapacity = 0;
	v.head = a;

	return &v;
}

func (v Vector) size() int {
	return v.numberOfItems;
}

func (v Vector) capacity() int {
	return v.arrayCapacity;
}

func (v Vector) is_empty() bool {
	return v.numberOfItems == 0;
}

func (v Vector) at(index int) (int, error) {
	if (index < 0 || index >= v.size()) {
		return -1, errors.New("Index out of bounds");
	}
	return v.head[index], nil;
}

//func (v Vector) push(item int) {
//	v.head[v.size()] = item;
//	v.numberOfItems++;
//}

func main() {
	v := NewVector();

	fmt.Println("Size of v: ", v.size());

	fmt.Println("Capacity of v: ", v.capacity());

	fmt.Println("Vector v is empty: ", v.is_empty());

	if r, e := v.at(0); e != nil {
		fmt.Println("Vector v returns error: ", e)
	} else {
		fmt.Println("Vector v at index 0: ", r)
	}
	//
	//v.push(1);
	//if r, e := v.at(0); e != nil {
	//	fmt.Println("Vector v returns error: ", e)
	//} else {
	//	fmt.Println("Vector v at index 0: ", r)
	//}

}
