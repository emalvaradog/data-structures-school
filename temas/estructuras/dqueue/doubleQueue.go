package dqueue

import (
	"fmt"
	"strings"
)

// DataDQueue public struct
type DataDQueue struct {
	Value int
}

// Node public struct
type Node struct {
	Data DataDQueue
	next *Node
}

// Queue public struct
type Queue struct {
	cursor int // -> Existent nodes
	head   *Node
	tail  *Node
}

// CreateDoubleQueue -> Create a new queue
func CreateDoubleQueue() *Queue {
	return new(Queue)
}

// PushBack -> Add element to end of queue
func (q *Queue) PushBack(data DataDQueue) {
	newNode := &Node{Data: data}

	if q.isEmpty(){
		q.head = newNode
		q.tail = newNode	
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}

	q.cursor ++

}

// PushFront -> Add element to front of queue
func (q *Queue) PushFront(data DataDQueue) {
	newNode := &Node{Data: data}
	if q.isEmpty() {
		q.head = newNode
		q.tail = newNode
	} else {
		newNode.next = q.head
		q.head = newNode
	}

	q.cursor ++
}

// PopFront -> Remove element from front of queue
func (q *Queue) PopFront(){
	if q.isEmpty() {
		fmt.Println("Empty queue")
	} else if q.cursor == 1{
		q.head = nil
		q.tail = nil
		q.cursor = 0
	} else {
		q.cursor--
		q.head = q.head.next
	}
}

// PopBack -> Remove element from end of queue
func(q *Queue) PopBack(){
	if q.isEmpty() {
		fmt.Println("Empty queue")
	} else if q.cursor == 1 {
		q.head = nil
		q.tail = nil
		q.cursor = 0

		it := q.head
		for ; it.next != q.tail; it = it.next {
		}
		q.tail = it
		q.tail.next = nil
		q.cursor--
	}
}


func (q *Queue) isEmpty() bool {
	if q.cursor == 0 && q.head == nil  && q.tail == nil{
		return true
	}
	return false
}


func (q Queue) String() string {
	if q.cursor == 0{
		return fmt.Sprint("Empty queue")
	}

	sb := strings.Builder{}

	for it := q.head; it != nil; it = it.next{
		sb.WriteString(fmt.Sprintf("%d", it.Data.Value))
	}

	return sb.String()
}
