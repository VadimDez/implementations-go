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
	var a = make([]int, 0, 8);
	var v = Vector{};

	v.numberOfItems = 0;
	v.arrayCapacity = 8;
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
	if (index < 0 || index >= v.numberOfItems) {
		return -1, errors.New("Index out of bounds");
	}
	return v.head[index], nil;
}

func (v *Vector) push(item int) {
	if (v.size() == v.capacity()) {
		v.resize(v.capacity() * 2)
	}

	v.head = append(v.head, item);
	v.numberOfItems++
}

func (v *Vector) resize(newCapacity int) {
	var a = make([]int, v.size(), newCapacity)
	var i int
	for i = 0; i < v.size(); i++ {
		a[i] = v.head[i]
	}
	v.head = a
	v.arrayCapacity = newCapacity
}

func (v *Vector) insert(index int, item int) {
	var i int

	for i = v.size(); i > index; i-- {
		v.head[i] = v.head[i - 1];
	}

	v.head[index] = item;
	v.numberOfItems++;
}

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

	v.push(1);

	if r, e := v.at(0); e != nil {
		fmt.Println("Vector v returns error: ", e)
	} else {
		fmt.Println("Vector v at index 0: ", r)
	}

	v.push(2);
	v.push(3);
	v.push(4);
	v.push(5);
	v.push(6);
	v.push(6);
	v.push(8);
	fmt.Println("vector before resize: ", v.head)
	v.push(9);
	fmt.Println("vector after resize: ", v.head)

	v.insert(1, 100);
	fmt.Println("vector after insert: ", v.head)
}
