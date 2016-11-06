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
	var a = make([]int, 8, 8);
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

	v.head[v.numberOfItems] = item;
	v.numberOfItems++
}

func (v *Vector) resize(newCapacity int) {
	var a = make([]int, newCapacity, newCapacity)
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

func (v *Vector) prepend(item int) {
	v.insert(0, item)
}

func (v *Vector) pop() int {
	v.numberOfItems--;

	if v.size() == v.capacity() / 4 {
		v.resize(v.capacity() / 2)
	}

	return v.head[v.numberOfItems];
}

func (v *Vector) deleteAt(index int) {
	var i int
	v.numberOfItems--

	for i = index; i < v.size(); i++ {
		v.head[i] = v.head[i + 1];
	}
}

func (v *Vector) remove(item int) {
	index := v.find(item)

	for index != -1 {
		v.deleteAt(index)

		index = v.find(item)
	}
}

func (v *Vector) find(item int) int {
	i := 0
	found := false

	for i < v.size() && !found {
		if v.head[i] == item {
			found = true
		}
		i++
	}

	if found {
		return i - 1
	}

	return -1
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
	v.push(7);
	v.push(8);
	fmt.Println("vector before resize: ")
	printArray(v)
	v.push(9);
	fmt.Println("vector after resize: ")
	printArray(v)

	v.insert(1, 100);
	fmt.Println("vector after insert: ")
	printArray(v)

	v.prepend(-1);
	fmt.Println("vector after prepend: ")
	printArray(v)

	fmt.Println("vector pop: ", v.pop())
	fmt.Println("vector after pop: ")
	printArray(v)

	v.deleteAt(2)
	fmt.Println("vector after delete item at index 2: ")
	printArray(v)

	fmt.Println("Find 8 in vector: ", v.find(8))

	v.push(8)
	printArray(v)
	v.remove(8)
	printArray(v)
}

func printArray(v *Vector) {
	var i int

	for i = 0; i < v.size(); i++ {
		fmt.Print(v.head[i], " ")
	}
	fmt.Print("\n")
}