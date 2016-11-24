package main

import (
	"math"
	"fmt"
	"errors"
)

type Item struct {
	key int
	value int
	isDeleted bool
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
	found := false
	end := false
	i := 0

	for i < this.size && !found && !end {
		position := (hash + i) % this.size // circular

		if this.isNull(this.array[position]) { // do not proceed
			end = true
		} else {
			if !this.array[position].isDeleted && this.array[position].key == hash {
				found = true
			}

			i++
		}
	}

	return found
}

func (_ HashTable) isNull(item Item) bool {
	return Item{} == item
}

func (this *HashTable) add(key int, value int) error {
	hash := this.hash(key, this.size)
	added := false
	i := 0

	for i < this.size && !added {
		position := (hash + i) % this.size // make it circular

		if this.isNull(this.array[position]) || this.array[position].isDeleted {
			this.array[position] = Item{key, value, false}
			this.count++
			added = true
		} else if this.array[position].key == key {
			this.array[position].value = value
			added = true
		}

		i++
	}

	if !added {
		return errors.New("Error: item wasn't added")
	}

	if this.count >= this.size * 3/4 {
		this.resize(this.size * 2)
	}

	return nil
}


func (this HashTable) get(key int) (int, error) {
	hash := this.hash(key, this.size)
	found, position := this.findWith(hash, key)

	if !found {
		return 0, errors.New("Key is not present in hash table")
	}

	return this.array[position].value, nil
}

func (this *HashTable) remove(key int) error {
	hash := this.hash(key, this.size)
	found, position := this.findWith(hash, key)

	if !found {
		return errors.New("Key is not present in hash table")
	}

	this.array[position].isDeleted = true
	this.array[position].value = -1
	this.array[position].key = -1

	this.count--

	if this.count <= this.size / 2 {
		this.resize(this.size / 2)
	}

	return nil
}

// resize and rehash all items
func (this *HashTable) resize(size int) {
	newTable := HashTable{}
	newTable.size = size
	newTable.count = 0
	newTable.array = make([]Item, size)

	for i := 0; i < this.size; i++ {
		item := this.array[i]

		if !this.isNull(item) && !item.isDeleted {
			newTable.add(item.key, item.value)
		}
	}

	this.array = newTable.array
	this.size = newTable.size
}

// Get position of item by key if found
func (this HashTable) findWith(hash int, key int) (bool, int) {
	i := 0
	found := false
	end := false
	var position int

	for i < this.size && !found && !end {
		position = (hash + i) % this.size

		if this.isNull(this.array[position]) {
			end = true
		} else {
			if !this.array[position].isDeleted && this.array[position].key == key {
				found = true
			}
		}

		i++
	}

	return found, position
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

	fmt.Println("hash of k=1, m=8: ", table.hash(1, 8));

	fmt.Println("Does key 1 exists: ", table.exists(1))

	if e := table.add(1, 100); e != nil {
		fmt.Println("add error: ", e)
	}

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

	if e := table.add(1, 1); e != nil {
		fmt.Println("add error: ", e)
	}
	if e := table.add(9, 9); e != nil {
		fmt.Println("add error: ", e)
	}

	table.add(9, 9)
	table.add(1, 9)
	table.add(2, 9)
	table.add(3, 9)
	table.add(4, 9)
	table.add(5, 9)
	table.add(6, 9)
	table.add(7, 9)
	table.add(8, 9)
	if err := table.add(10, 9); err != nil {
		fmt.Println("Error: ", err)
	}



	v, e := table.get(1)
	fmt.Println("Get value at key: 1 :", v, "Error: ", e)
	v, e = table.get(9)
	fmt.Println("Get value at key: 9 :", v, "Error: ", e)
	fmt.Println("length: ", table.count)
	fmt.Println("length: ", table.size)
}
