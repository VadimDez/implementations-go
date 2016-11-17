package main

import (
	"fmt"
	"errors"
)

type Queue struct {
	// queue - fix-sized array
	array []int
	// head index
	head int
	// tail index
	tail int
	// max size
	size int
}

// Check if queue is empty
func (q Queue) isEmpty() bool {
	return q.head == q.tail
}

func (q *Queue) enqueue(value int) error {
	if q.isFull() {
		return errors.New("Queue is full")
	}

	q.array[q.tail] = value
	q.tail = q.getIndex(q.tail + 1)

	return nil
}

// dequeue item
func (q *Queue) dequeue() (int, error) {
	if q.isEmpty() {
		return -1, errors.New("Queue is empty")
	}

	value := q.array[q.head]
	q.head = q.getIndex(q.head + 1)
	return value, nil
}

// Adjust index
func (q Queue) getIndex(index int) int {
	if index < 0 {
		return q.size + index
	}

	if index >= q.size {
		return index - q.size
	}

	return index
}

// check if queue is full
func (q Queue) isFull() bool {
	return q.getIndex(q.tail + 1) == q.head
}

/**
Create queue
 */
func newQueue(size int) *Queue {
	var q = Queue{}

	q.array = make([]int, size)
	q.head = 0
	q.tail = 0
	q.size = size

	return &q
}

func main() {
	q := newQueue(5)

	fmt.Println("Is queue empty: ", q.isEmpty())

	q.enqueue(1)
	fmt.Println("Is queue empty after enqueue: ", q.isEmpty())

	if v, e := q.dequeue(); e == nil {
		fmt.Println("Dequeue: ", v)
	}
	fmt.Println("Is queue empty after enqueue: ", q.isEmpty())


	q.enqueue(1)
	q.enqueue(2)
	q.enqueue(3)
	q.enqueue(4)
	q.enqueue(5)
	if e := q.enqueue(6); e != nil {
		fmt.Println("Enqueue error: ", e)
	}

	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.dequeue()
	q.dequeue()

	if _, e := q.dequeue(); e != nil {
		fmt.Println("Dequeue error: ", e)
	}
}
