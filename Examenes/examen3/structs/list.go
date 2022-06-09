package structs

import (
	"fmt"
	"os"
)

// DataList struct
type DataList struct {
	Valor string
	Freq int
}

// NodeList struct
type NodeList struct {
	Info Tree
	Next* NodeList
}

// List struct
type List struct {
	Head* NodeList
	Len int
	Order int
}

// CreateList public function
func CreateList() *List {
	return new(List)
}

// AddRight public method
func (l* List) AddRight(data DataList) {
	node := &NodeList{Info: Tree{Data: TData{Char: data.Valor, Freq: data.Freq}}}
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
	node := &NodeList{Info: Tree{Data: TData{Char: data.Valor, Freq: data.Freq}}}
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
	node := &NodeList{Info: Tree{Data: TData{Char: data.Valor, Freq: data.Freq}}}
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
	fmt.Printf("\n")
	fmt.Printf("Lista | Len %d \n", l.Len)
	var i int

	n := l.Head

	if n == nil {
		fmt.Println("Lista vacía")

	} else {
		for ; n != nil ; n = n.Next {
			fmt.Printf("Value: %s | Freq %d | Left: %+v | Right: %+v \n", n.Info.Data.Char, n.Info.Data.Freq, n.Info.Left, n.Info.Right)
			i++
		}
		// fmt.Printf("DataList: %d | Pos: %d \n", n.Info.Valor, i)
	}
}


// Contains returns true if list contains certain strng
func (l* List) Contains(s string) bool{
	n := l.Head

	if n == nil {
		return false
	}

	for ; n != nil ; n = n.Next {
		if n.Info.Data.Char == s {
			return true
		}
	}

	return false
}

// IncreaseFreq increases freq given a match
func (l *List) IncreaseFreq(s string) {	
	n := l.Head

	if n == nil {
		return
	}

	for ; n != nil; n = n.Next {
		if n.Info.Data.Char == s {
			n.Info.Data.Freq++
		}
	}

}

// OrderList orders list
func (l *List)OrderList() {
	var pivote, aux *NodeList
	var guardar Tree

	pivote = l.Head.Next

	for pivote != nil {
		aux = l.Head

		for aux != pivote {
			guardar = aux.Info

			if aux.Info.Data.Freq > pivote.Info.Data.Freq {

				aux.Info.Data.Freq = pivote.Info.Data.Freq
				aux.Info.Data.Char = pivote.Info.Data.Char
				aux.Info.Left = pivote.Info.Left
				aux.Info.Right = pivote.Info.Right

				pivote.Info.Data.Freq = guardar.Data.Freq
				pivote.Info.Data.Char = guardar.Data.Char
				pivote.Info.Left = guardar.Left
				pivote.Info.Right = guardar.Right

			}
			aux = aux.Next
		}
		pivote = pivote.Next
	}
}

// Huffman function
func (l *List) Huffman(charList[]TData) {
	for l.Len != 1 {
		l.OrderList()
		l.TreeFromNodes()
	}

	var lines []string


	for i := 0; i < len(charList); i++ {
		lines = append(lines, l.Head.Info.SearchNode(charList[i]))
	}

	for i:= 0; i<len(lines) ; i++ {
		fmt.Println(lines[i])
	}

  
  var file, err = os.OpenFile("CodificacionHuffman.txt", os.O_RDWR, 0644)
    
	if err != nil {
		fmt.Println("E")
	}

    defer file.Close()

		for i := 0; i < len(lines) ; i++{
			_, err = file.WriteString("lines")
		}
    

}

// TreeFromNodes function
func (l *List) TreeFromNodes() {
	
	first := &l.Head.Info
	second := &l.Head.Next.Info

	concat := fmt.Sprintf("%d%s", l.Len, "_")

	root := Tree{}
	root.Data.Char = concat
	root.Data.Freq = first.Data.Freq + second.Data.Freq / 2

	if l.Order == 0 {
		root.AddNode(first.Data)
		root.AddNode(second.Data)
	} else {
		root.AddTree(l.Head.Info)
		root.AddTree(l.Head.Next.Info)
	}

	root.Data.Freq = first.Data.Freq + second.Data.Freq

	l.Head.Info = root
	l.Head.Next = l.Head.Next.Next
	l.Len--	
	l.Order++
}
