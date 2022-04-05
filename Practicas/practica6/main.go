package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	a "practica6/structs"
	"strconv"
)

func balancear() {
	sD := a.CrearStack()

	file, err := os.Open("balanceo.txt")
	if err != nil {
		log.Fatal(err)
	}
	fScan := bufio.NewScanner(file)
	fScan.Split(bufio.ScanLines)
	var lines []string
	for fScan.Scan() {
		lines = append(lines, fScan.Text())
	}

	file.Close()

	for i := 0; i < len(lines); i++ {
		for _, l := range lines[i] {
			if string(l) == "(" {
				n := a.Nodo{Value: "("}
				sD.Push(n)
			} else if string(l) == ")"{
				e := a.Nodo{}
				sD.Pop(&e)
			}
			
		}
		fmt.Println(lines[i], sD.IsEmpty())
		sD.Empty()
	}
}

func calcular() {
	
	// Read file and append lines
	file, err := os.Open("expresiones.txt")
	if err != nil {
		log.Fatal(err)
	}

	fScan := bufio.NewScanner(file)
	fScan.Split(bufio.ScanLines)
	var lines []string

	for fScan.Scan() {
		lines = append(lines, fScan.Text())
	}
	
	file.Close()

	// Create dynamic stack
	sD := a.CrearDstack()

	// Funcional variable
	ops := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	// Load data
	for i := 0; i < len(lines); i++ {
		for _, car := range lines[i] {
			i1 := a.Node{}
			i2 := a.Node{}
			if(string(car) == "*" || string(car) == "+"){
				if !sD.IsEmpty() {
					sD.Pop(&i1)
					sD.Pop(&i2)
					res := a.Node{Value:ops[string(car)](i1.Value, i2.Value)}
					sD.Push(res)
				} else {
					continue
				}
			} else {
				va, _ := strconv.Atoi(string(car))
				v := a.Node{Value: va}
				sD.Push(v)
			}
		}
		fmt.Println(lines[i], "=", sD.ConsultarTop())
		sD.Empty()
	}	
}

func main() {
	fmt.Println("-- Evaluación de paréntesis --")
	balancear()	
	fmt.Printf("\n \n")
	fmt.Println("-- Operaciones con expresiones posfijas --")
	calcular()
}