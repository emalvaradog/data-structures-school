package structs

import "fmt"

// Node public struct
type Node struct {
	Value int
}

// Dstack public struct
type Dstack struct {
	Data []Node
	Top, Err, Cap int
}

// CrearDstack - Create Dstack
func CrearDstack() *Dstack {
	var st Dstack
	st.Top = -1
	st.Err = 0
	st.Cap = 0
	return &st
}

func (s Dstack) String() string {
	return fmt.Sprintf("Top: %d | Err: %d | Data: %v\n", s.Top + 1, s.Err, s.Data)
}

// ConsultarTop - Get top element
func (s *Dstack) ConsultarTop() int {
	return s.Data[s.Top].Value
}

// IsEmpty - Check if empty
func (s *Dstack) IsEmpty() bool {
	if s.Top == -1 {
		return true
	}
	return false
}

// IsFull - Check if full
func (s *Dstack) IsFull() bool {
	if s.Top == s.Cap - 1 {
		return true
	} 
	return false
}

// Push - Append element to stack
func (s *Dstack) Push(data Node) {
	if s.IsFull() {
		(*s).Data = append((*s).Data, data)
		(*s).Top++
		(*s).Cap++
		(*s).Err = 0
	} else {
		(*s).Data = append((*s).Data, data)
		(*s).Top++
		(*s).Cap++
	}
}

// Pop - Remove last element
func (s *Dstack) Pop(extract *Node) {
	if !s.IsEmpty() {
		*extract = s.Data[s.Top]
		s.Data = s.Data[:s.Top]
		(*s).Top--
		(*s).Err = 0
	} else {
		fmt.Println("Pila vacía")
		(*s).Err = -2
	}
}

// Empty - Free stack
func (s *Dstack) Empty() {
	var aux Node
	for !s.IsEmpty() {
		s.Pop(&aux)
		// fmt.Println("Valor extraido es ", aux.Value)
	}
	// fmt.Println("Empty Dstack")
}
