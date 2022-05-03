package main 

import (
	"practica8/structs"
	"fmt"
)

func main() {
	l := structs.CreateList()
	fmt.Println(l)
	var opc int
	for opc != 7 {
		fmt.Println("Opcion")
		fmt.Println("0 - Ingresar nodo por frente")
		fmt.Println("1 - Ingresar nodo por medio")
		fmt.Println("2 - Ingresar nodo al final")
		fmt.Println("3 - Buscar elemento en la lista")
		fmt.Println("4 - Eliminar un nodo")
		fmt.Println("5 - Imprimir nodos")
		fmt.Println("6 - Eliminar toda la lista")
		fmt.Println("7 - Salir")
		fmt.Scanl(&opc)
		switch opc {
			case 0:
				data := structs.DataList
				var value int
				fmt.Println("Ingresar valor")
				fmt.Scanln(&data)
				data.Valor = data
				l.AddLeft(data)

			case 1:
				data := structs.DataList
				half = l.Len / 2
				var value int
				fmt.Println("Ingresar valor")
				fmt.Scanln(&data)
				data.Valor = data
				l.AddNPos(data, half)
			
			case 2:
				data := structs.DataList
				var value int
				fmt.Println("Ingresar valor")
				fmt.Scanln(&data)
				data.Valor = data
				l.AddLeft(data)

			case 3: 
				var pos int
				fmt.Println("Ingresa valor")
				fmt.Scanln(&pos)
				a := l.SearchPos(pos)
				fmt.Println(a)

			case 4:
				var pos int
				fmt.Println("Ingresa valor")
				fmt.Scanln(&pos)
				l.DeleteNode(pos)
			
			case 5:
				l.PrintList()

			case 6:
				fmt.Println("Delete list")
				l.DeleteList()
				
		}
	}
}