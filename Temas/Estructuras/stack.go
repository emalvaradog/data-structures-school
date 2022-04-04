package estructuras

import "fmt"

type node struct {
	value int
}

type stack struct {
	elements [5]node
	top, err int
}

func crearTDA() *stack {
	var st stack
	st.top = -1
	st.err = 0
	return &st
}

func (s stack) String() string {
	return fmt.Sprintf("Top: %d | err: %d | elements: %v\n", s.top + 1, s.err, s.elements)
}

func (s *stack) consularTop() {
	fmt.Println("TOP:", s.top)
}

func (s *stack) isEmpty() bool {
	if s.top == -1 {
		return true
	}
	return false
}

func (s *stack) isFull() bool {
	if s.top == len(s.elements) - 1 {
		return true
	} 
	return false
}

func (s *stack) push(data node) {
	if !s.isFull() {
		s.top++
		s.elements[s.top] = data
		s.err = 0
	} else {
		fmt.Println("Pila llena")
		s.err = -1
	}
}

func (s *stack) pop(extract *node) {
	if !s.isEmpty() {
		*extract = s.elements[s.top]
		s.top--
		s.err = 0
	} else {
		fmt.Println("Pila vacía")
		s.err = -2
	}
}

func (s *stack) empty(st *stack, error *int) {
	var aux node
	for !s.isEmpty() {
		s.pop(&aux)
		fmt.Println("Valor extraido es ", aux.value)
	}
	fmt.Println("Empty stack")
}

func main() {
	s := crearTDA()
	for i := 0; i < 5; i++ {
		fmt.Println("Ingresar valor")
		var in node
		fmt.Scanln(&in.value)
		s.push(in)
	}
	fmt.Println(s)
}
