package structs

import (
	"fmt"
	"strings"
	"time"
)

// Data public struct
type Data struct {
	Value string
	Title string
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

// CreateDoubleQueue -> Create a new queue
func CreateDoubleQueue() *Queue {
	q := new(Queue)
	q.cursor = 0
	return q
}

// PushBack -> Add element to end of queue
func (q *Queue) PushBack(data Data) {
	newNode := &Node{Data: data}

	if q.IsEmpty(){
		q.head = newNode
		q.tail = newNode	
	} else {
		q.tail.next = newNode
		q.tail = newNode
	}

	q.cursor ++

}

// PushFront -> Add element to front of queue
func (q *Queue) PushFront(data Data) {
	newNode := &Node{Data: data}

	if q.IsEmpty() {
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
	if q.IsEmpty() {
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

// GetHead -> Get first element of queue
func (q *Queue) GetHead() string {
	return (fmt.Sprintf("%s | %s", q.head.Data.Title, q.head.Data.Value))
}

// GetTail -> Get last element of queue
func (q *Queue) GetTail() string {
	return q.tail.Data.Value
}

// PopBack -> Remove element from end of queue
func(q *Queue) PopBack(){
	if q.IsEmpty() {
		fmt.Println("Empty queue")
	} else if q.cursor == 1 {
		q.head = nil
		q.tail = nil
		q.cursor = 0
	} else {
		it := q.head
		for ; it.next != q.tail; it = it.next {
		}
		q.tail = it
		q.tail.next = nil
		q.cursor--
	}
}

// IsEmpty -> Validate if queue is empty
func (q *Queue) IsEmpty() bool {
	if q.cursor == 0 && q.head == nil  && q.tail == nil{
		return true
	}
	return false
}

// Clear -> Remove all elements from the queue
func (q *Queue) Clear() {
	if q.IsEmpty() {
		fmt.Println("Queue is already empty")
	} else {
		q.head = nil
		q.tail = nil
		q.cursor = 0
		fmt.Println("Cleared queue succesfully")
	}
}

// Send -> Simulate file sending from queue
func (q *Queue) Send() {
	for it := q.head; it != nil; it = it.next {
		fmt.Println("Sending ", it.Data.Title)
		fmt.Println(it.Data.Value)
		time.Sleep(time.Second * 3)
		fmt.Println("Sent succesfully")
		fmt.Printf("\n")
	}
	fmt.Println("Nothing to send")
}

func (q Queue) String() string {
	if q.cursor == 0{
		return fmt.Sprint("Empty queue")
	}

	sb := strings.Builder{}

	for it := q.head; it != nil; it = it.next{
		sb.WriteString(fmt.Sprintf("File: %s | Content: %s \n", it.Data.Title, it.Data.Value))
	}

	return sb.String()
}
