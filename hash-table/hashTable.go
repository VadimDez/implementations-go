package main

import (
	"math"
	"fmt"
	"errors"
)

type Item struct {
	key int
	value int
	deleted bool
}

type HashTable struct {
	array []Item
	size int
	count int
}

func (this HashTable) hash(k int, m int) int {
	return int(math.Abs(float64(k))) % m
}

func (this HashTable) exists(key int) bool {
	hash := this.hash(key, this.size)

	// no collisions for now
	return !this.isNull(this.array[hash])
}

func (_ HashTable) isNull(item Item) bool {
	return Item{} == item
}

func (this *HashTable) add(key int, value int) {
	hash := this.hash(key, this.size)

	this.array[hash] = Item{key, value, false}

	this.count++
}

func (this HashTable) get(key int) (int, error) {
	hash := this.hash(key, this.size)

	if this.isNull(this.array[hash]) {
		return 0, errors.New("Key is not present in hash table")
	}

	return this.array[hash].value, nil
}

func (this *HashTable) remove(key int) error {
	hash := this.hash(key, this.size)

	if this.isNull(this.array[hash]) {
		return errors.New("Key is not present in hash table")
	}

	this.array[hash] = Item{}
	this.count--

	return nil
}

func NewHashTable() *HashTable {
	table := HashTable{}

	table.size = 8
	table.count = 0
	table.array = make([]Item, table.size)

	return &table
}

func main() {
	table := NewHashTable()

	fmt.Println(table.hash(1, 8));

	fmt.Println("Does key 1 exists: ", table.exists(1))

	table.add(1, 100)

	fmt.Println("Does key 1 exists now: ", table.exists(1))

	fmt.Println("Get value at key: 1")

	v, _ := table.get(1)

	fmt.Println(v)


	fmt.Println("Get value at key that doesn't exist yet: 2")

	if _, e := table.get(2); e != nil {
		fmt.Println("error: ", e)
	}

	fmt.Println("Remote item with key: 1")
	if e := table.remove(1); e == nil {
		fmt.Println("Removed")
	} else {
		fmt.Println("Error during removing: ", e)
	}
	fmt.Println("Does key 1 still exists: ", table.exists(1))
}
