package list

type Dato struct {
	valor int
}

type Nodo struct {
	info Dato
	next* Nodo
	pos int
}

type List struct {
	head* Nodo
}

func createList() *List {
	var l List
	return &l
}

func (l* List) addRight(data Data) {
	node := Nodo{info: data}
	var pos int
	n := l.head
	for ; n.next != nil; n = n.next {	
		pos++
	}
	node.pos = pos
	n.next = node
}

func (l* List) addLeft(data Dato) {
	node := Nodo{info: data, pos: 0}
	node.next = l.head
	l.head = node
}

func (l* List) addNPos(data Dato, pos int) {
	var counter int
	n := l.head

	for ; counter != pos-1 ; counter++{
		n = n.next
	}

	node := Nodo{info: data}
	node.next = n.next
	n.next = node
}


func (l* List) imprimirList() {
	n := l.head
	var i int
	for ; n.next != nil ; n = n.next {
		i++
		fmr.Printf("Dato: %d | Pos: %d", n.info.valor, i)
	}
}