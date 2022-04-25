package queue

import (
	"fmt"
	"strings"
)

// Data public struct
type Data struct {
	Value int
}

// Node public struct
type Node struct {
	Data Data
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
	q := new(Queue)
	q.cursor = 0
	return q
}

// Enqueue -> Add new element to queue
func (q *Queue) Enqueue(data Data) {
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
func (q *Queue) Dequeue(){
	if q.isEmpty() {
		fmt.Println("Queue is empty")
	} else if q.cursor == 1 {
		q.head = nil
		q.tail = nil
		q.cursor = 0
	} else {
		q.head = q.head.next
		q.cursor --
	}
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
		sb.WriteString(fmt.Sprintf("%d", it.Data.Value))
	}

	return sb.String()
}
