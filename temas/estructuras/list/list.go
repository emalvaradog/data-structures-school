package list

import "fmt"

// Dato struct
type Dato struct {
	Valor int
}

// Nodo struct
type Nodo struct {
	Info Dato
	Next* Nodo
}

// List struct
type List struct {
	Head* Nodo
	Len int
}

// CreateList public function
func CreateList() *List {
	return new(List)
}

// AddRight public method
func (l* List) AddRight(data Dato) {
	node := &Nodo{Info: data}
	n := l.Head
	if l.Head == nil {
		l.Head = node
	} else {
		for ; n.Next != nil; n = n.Next {	
		}
		n.Next = node
	}

	l.Len++
}

// AddLeft public method
func (l* List) AddLeft(data Dato) {
	node := &Nodo{Info: data}
	if l.Head == nil {
		l.Head = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
	
	l.Len++
}

// AddNPos public metod
func (l* List) AddNPos(data Dato, pos int) {
	var counter int
	node := &Nodo{Info: data}	
	n := l.Head

	if l.Head == nil {
		l.Head = node
	} else if pos == 0 {
		node.Next = l.Head
		l.Head = node 
	}	else if pos > l.Len {
		fmt.Printf("Position not valid, you want to inset into %d, but list only has %d \n", pos, l.Len)
	} else {
		for ; counter != pos-1 ; counter++{
			n = n.Next
		}
		node.Next = n.Next
		n.Next = node
	}

	l.Len++
}

// DeleteList public method
func (l* List) DeleteList() {
	l.Head = nil
}

// PrintList public method
func (l* List) PrintList() {
	var i int

	n := l.Head

	if n == nil {
		fmt.Println("Lista vacía")

	} else {
		for ; n != nil ; n = n.Next {
			fmt.Printf("Dato: %d | Pos: %d \n", n.Info.Valor, i)
			i++
		}
		// fmt.Printf("Dato: %d | Pos: %d \n", n.Info.Valor, i)
	}
}

