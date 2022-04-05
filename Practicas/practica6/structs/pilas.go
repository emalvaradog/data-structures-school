package structs

import "fmt"

// Nodo public struct
type Nodo struct {
	Value string
}

// Stack public struct
type Stack struct {
	Data [50]Nodo
	Top, Err int
}

// CrearStack public func
func CrearStack() *Stack {
	var st Stack
	st.Top = -1
	st.Err = 0
	return &st
}

func (s Stack) String() string {
	return fmt.Sprintf("Top: %d | Err: %d | Data: %v\n", s.Top, s.Err, s.Data)
}

// ConsultarTop - Public method
func (s *Stack) ConsultarTop() {
	fmt.Println("Top:", s.Top)
}

// IsEmpty - Public method
func (s *Stack) IsEmpty() bool {
	if s.Top == -1 {
		return true
	}
	return false
}

// IsFull - Public method
func (s *Stack) IsFull() bool {
	if s.Top == len(s.Data) - 1 {
		return true
	} 
	return false
}

// Push - Public method
func (s *Stack) Push(data Nodo) {
	if !s.IsFull() {
		(*s).Top++
		(*s).Data[s.Top] = data
		(*s).Err = 0
	} else {
		fmt.Println("Pila llena")
		(*s).Err = -1
	}
}

// Pop - public method
func (s *Stack) Pop(extract *Nodo) {
	if !s.IsEmpty() {
		*extract = s.Data[s.Top]
		s.Top--
		s.Err = 0
	} else {
		fmt.Println("Pila vacía")
		s.Err = -2
	}
}

// Empty - Public method
func (s *Stack) Empty() {
	var aux Nodo
	for !s.IsEmpty() {
		s.Pop(&aux)
	}
}
