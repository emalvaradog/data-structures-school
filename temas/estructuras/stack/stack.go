package stack

import "fmt"

// NodeStack public struct
type NodeStack struct {
	Value int
}

// Stack public struct
type Stack struct {
	Data [5]NodeStack
	Top, Err int
}

func crearStack() *Stack {
	var st Stack
	st.Top = -1
	st.Err = 0
	return &st
}

func (s Stack) String() string {
	return fmt.Sprintf("Top: %d | Err: %d | Data: %v\n", s.Top + 1, s.Err, s.Data)
}

func (s *Stack) consularTop() {
	fmt.Println("Top:", s.Top)
}

func (s *Stack) isEmpty() bool {
	if s.Top == -1 {
		return true
	}
	return false
}

func (s *Stack) isFull() bool {
	if s.Top == len(s.Data) - 1 {
		return true
	} 
	return false
}

func (s *Stack) push(data NodeStack) {
	if !s.isFull() {
		(*s).Top++
		(*s).Data[s.Top] = data
		(*s).Err = 0
	} else {
		fmt.Println("Pila llena")
		(*s).Err = -1
	}
}

func (s *Stack) pop(extract *NodeStack) {
	if !s.isEmpty() {
		*extract = s.Data[s.Top]
		s.Top--
		s.Err = 0
	} else {
		fmt.Println("Pila vacía")
		s.Err = -2
	}
}

func (s *Stack) empty(st *Stack, Error *int) {
	var aux NodeStack
	for !s.isEmpty() {
		s.pop(&aux)
		fmt.Println("Valor extraido es ", aux.Value)
	}
	fmt.Println("Empty Stack")
}

// func main() {
// 	s := crearTDA()
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("Ingresar valor")
// 		var in NodeStack
// 		fmt.Scanln(&in.Value)
// 		s.push(in)
// 	}
// 	fmt.Println(s)
// }
