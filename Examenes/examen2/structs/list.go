package structs

import "fmt"

// DataList struct
type DataList struct {
	Age int
	Name string
	Height float64
}

// NodeList struct
type NodeList struct {
	Info DataList
	Next* NodeList
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
		node.Next = n.Next
		n.Next = node
	}

	l.Len++
}

// ServeCustomers - Remove first two elements from list
func (l* List) ServeCustomers(r int) {
	fmt.Printf("Juego: %d \n", r)
	if l.Len == 0 {
		fmt.Println("No hay clientes")
		
	}	else if l.Len == 1 {
		fmt.Println("Ventanilla:")
		fmt.Printf("-- Nombre: %s | Altura: %0.2f \n", l.Head.Info.Name, l.Head.Info.Height)
		l.Head = nil
		l.Len--

	} else {
		fmt.Println("Ventanilla:")
		fmt.Printf("-- Nombre: %s | Altura: %0.2f \n", l.Head.Info.Name, l.Head.Info.Height)
		l.Head = l.Head.Next
		l.Len--

		fmt.Printf("-- Nombre: %s | Altura: %0.2f \n", l.Head.Info.Name, l.Head.Info.Height)
		l.Head = l.Head.Next
		l.Len--

		fmt.Printf("Pendientes: \n")
		l.PrintList()
	} 
	fmt.Println("")
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
			fmt.Printf("-- Nombre: %s | Altura: %0.2f \n", n.Info.Name, n.Info.Height)
			i++
		}
		// fmt.Printf("DataList: %d | Pos: %d \n", n.Info.Valor, i)
	}
}

