package main

import (
	"fmt"
	"topics/estructuras/list"
)

func main() {
	fmt.Println("List")
	d1 := list.Dato{Valor: 10}
	d2 := list.Dato{Valor: 5}
	d3 := list.Dato{Valor: 0}
	d4 := list.Dato{Valor: 50}
	l := list.CreateList()
	l.AddLeft(d1)
	l.AddLeft(d2)
	l.PrintList()
	fmt.Println("")
	l.AddNPos(d3, 2)
	l.AddNPos(d4, 0)
	l.PrintList()
}
