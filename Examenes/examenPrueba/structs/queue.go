package structs

import (
	"fmt"
	"strings"
)

// DataQueue public struct
type DataQueue struct {
	Name string
	Age int
	Height float64
}

// Node public struct
type Node struct {
	Data DataQueue
	next *Node
}

// Queue public struct
type Queue struct {
	cursor int // -> Existent nodes
	head   *Node
	tail  *Node
}

// CreateQueue -> Create a new queue
func CreateQueue() *Queue {
	return new(Queue)
}

// Enqueue -> Add new element to queue
func (q *Queue) Enqueue(data DataQueue) {
	newNode := &Node{Data: data}
	if q.isEmpty() {
		q.head = newNode
		q.tail = newNode
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}
	q.cursor ++
}

// Dequeue -> Remove first element from queue
func (q *Queue) Dequeue() DataQueue{
	var n *Node
	if q.isEmpty() {
		fmt.Println("Queue is empty")
	} else if q.cursor == 1 {
		n = q.head
		q.head = nil
		q.tail = nil
		q.cursor = 0
	} else {
		n = q.head
		q.head = q.head.next
		q.cursor --
	}
	
	return n.Data
}

// Clear -> Clear all queue from nodes
func (q *Queue) Clear() {
	if q.isEmpty() {
		fmt.Println("Queue is empty")
	} else {
		q.head = nil
		q.tail = nil
		q.cursor = 0
	}
}

func (q *Queue) isEmpty() bool {
	if q.head == nil && q.tail == nil && q.cursor == 0{
		return true
	} 
	return false
}

func (q Queue) String() string {
	if q.cursor == 0{
		return fmt.Sprint("Queue is empty")
	}

	sb := strings.Builder{}

	for it := q.head; it != nil; it = it.next{
		sb.WriteString(fmt.Sprintf("%s | %f.2 | %d \n", it.Data.Name, it.Data.Height, it.Data.Age))
	}

	return sb.String()
}
