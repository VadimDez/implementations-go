package main

import (
	"math"
	"fmt"
)

type Item struct {
	key int
	value int
	deleted bool
}

type HashTable struct {
	array []Item
	size int
}

func (this HashTable) hash(k int, m int) int {
	return int(math.Abs(float64(k))) % m
}

func (this HashTable) exists(key int) bool {
	hash := this.hash(key, this.size)

}

func NewHashTable() *HashTable {
	table := HashTable{}

	table.size = 8
	table.array = make([]Item, table.size)

	return &table
}

func main() {
	table := NewHashTable()

	fmt.Println(table.hash(1, 8));
}
