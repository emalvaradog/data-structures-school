
package structs

import "fmt"

// DataList struct
type DataList struct {
	Valor int
}

// NodeList struct
type NodeList struct {
	Info DataList
	Next* NodeList
	Prev* NodeList
}

// List struct
type List struct {
	Head* NodeList
	Len int
}

// CreateList public function
func CreateList() *List {
	return new(List)
}

// AddRight public method
func (l* List) AddRight(data DataList) {
	node := &NodeList{Info: data}
	n := l.Head
	if l.Head == nil {
		l.Head = node
	} else {
		for ; n.Next != nil; n = n.Next {	
		}
		node.Prev = node
		n.Next = node
	}
	l.Len++
}

// AddLeft public method
func (l* List) AddLeft(data DataList) {
	node := &NodeList{Info: data}
	if l.Head == nil {
		l.Head = node
	} else {
		node.Next = l.Head
		l.Head = node
	}
	l.Len++
}

// AddNPos public metod
func (l* List) AddNPos(data DataList, pos int) {
	var counter int
	node := &NodeList{Info: data}	
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
		node.Prev = n
		node.Next = n.Next
		n.Next = node
	}
	l.Len++
}

func (l* List) SearchPos(pos int) NodeList{
	var counter int
	n := l.Head
	for ; counter != pos; counter++ {
		n = n.Next
	}

	return n
}

// DeleteList public method
func (l* List) DeleteList() {
	l.Head = nil
}

func (l* List) DeleteNode(pos int) {
	var counter int
	n := l.Head
	aux := new(NodeList)

	if pos == 0 {
		aux = n.next
		aux.Prev = nil
		n = aux	
	} else {
		for ; counter != pos ; counter++ {
			n = n.Next
		}
		n.Prev.Next = n.Next
		n.Next.Prev = n.Prev
	}
}

// PrintList public method
func (l* List) PrintList() {
	var i int
	n := l.Head
	if n == nil {
		fmt.Println("Lista vacía")
	} else {
		for ; n != nil ; n = n.Next {
			fmt.Printf("DataList: %d | Pos: %d \n", n.Info.Valor, i)
			i++
		}
		// fmt.Printf("DataList: %d | Pos: %d \n", n.Info.Valor, i)
	}
}
