package main

import (
	"examenprueba/structs"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strconv"
	"time"
)

func addClient(q* structs.Queue) {
	d := structs.DataQueue{}
	rand.Seed(time.Now().UnixNano())
	max := 1.0
	min := 2.0
	fmt.Println("Nombre")
	fmt.Scanln(&d.Name)
	fmt.Println("Edad")
	fmt.Scanln(&d.Age)
	 d.Height = min + rand.Float64() * (max - min)
	fmt.Println("Altura")
	fmt.Printf("%0.2f \n", d.Height)
	q.Enqueue(d)
	fmt.Println("")
}

func serveClients(l1, l2, l3, l4, l5 *structs.List) {
	counter := 0
	list := []int{l1.Len, l2.Len, l3.Len, l4.Len, l5.Len}
	sort.Ints(list)
	turn := 1
	var limit int

	if list[4] == 1 || list[4] == 2{
		limit = 1
	} else if list[4] == 3 {
		limit = 2
	} else {		
		limit = list[4] - 2
	}

	for; counter < limit; counter ++ {
		fmt.Printf("Turno %d \n \n", turn)
		l1.ServeCustomers(1)
		l2.ServeCustomers(2)
		l3.ServeCustomers(3)
		l4.ServeCustomers(4)
		l5.ServeCustomers(5)
		turn ++
		fmt.Printf("\n\n")
	}
}

func main() {

	fmt.Println("WELCOME")
	l1 := structs.CreateList()
	l2 := structs.CreateList()
	l3 := structs.CreateList()
	l4 := structs.CreateList()
	l5 := structs.CreateList()
	q := structs.CreateQueue()

	args, _ := strconv.Atoi(os.Args[1])
	for i := 0; i < args ; i++{
		addClient(q)
	}

	for i := 0; i < args; i++ {
		cq := q.Dequeue()
		cl := structs.DataList{Name: cq.Name, Age: cq.Age, Height: cq.Height}
		if cq.Height > 1.0 && cq.Height < 1.2 {
			l1.AddRight(cl)
		} else if cq.Height > 1.2 && cq.Height < 1.7 {
			l3.AddRight(cl)
		} else if cq.Height > 1.7 && cq.Height <= 2.0 {
			l5.AddRight(cl)
		} else if cq.Height == 1.2 || cq.Height == 1.7 {
			l4.AddRight(cl)
		} else {
			l2.AddRight(cl)
		}
	}

	serveClients(l1, l2, l3, l4, l5)

}