package main

import (
	"fmt"
	"topics/estructuras/tree"
)

func main() {
	t1 := tree.Tree{}
	t2 := tree.Tree{}
	
	var opc int
	for opc != 8 {
		fmt.Println("Menu")
		fmt.Println("1. AddNode")
		fmt.Println("2. SearchNode")
		fmt.Println("3. DeleteNode")
		fmt.Println("4. Depth")
		fmt.Println("5. In-Order")
		fmt.Println("6. Pre-Order")
		fmt.Println("7. JoinTree")
		fmt.Println("8. Leave")
		fmt.Scanf("%d", &opc)

		switch opc {
			case 1:
				data := tree.TData{}
				fmt.Println("Enter int value")
				fmt.Scanf("%d", &data.Value)
				
				var opc int
				fmt.Println("1. T1 | 2.T2")
				fmt.Scanf("%d", &opc)
				if opc == 1{
					t1.AddNode(data)
				} else if opc == 2 {
					t2.AddNode(data)
				}

				fmt.Println("Node Added")

			case 2:
				data := tree.TData{}
				fmt.Println("Enter int value")
				fmt.Scanf("%d", &data.Value)
				var res *tree.Tree

				var opc int
				fmt.Println("1. T1 | 2.T2")
				fmt.Scanf("%d", &opc)
				if opc == 1{
					res = t1.SearchNode(data)
				} else if opc == 2 {
					res = t2.SearchNode(data)
				}
				
				if res == nil {
					fmt.Println("Node not found")
				}

				fmt.Println("Node found")
				fmt.Printf("Value: %d | Left: %d | Right: %d \n", res.Data.Value, res.Left.Data.Value, res.Right.Data.Value)

			case 3:	
				data := tree.TData{}
				fmt.Println("Enter int value")
				fmt.Scanf("%d", &data.Value)

				var opc int
				fmt.Println("1. T1 | 2.T2")
				fmt.Scanf("%d", &opc)
				if opc == 1{
					t1.DeleteNode(data)
				} else if opc == 2 {
					t2.DeleteNode(data)
				}
				
				fmt.Println("Node deleted")

			case 4:

				var opc int
				fmt.Println("1. T1 | 2.T2")
				fmt.Scanf("%d", &opc)
				if opc == 1{
					fmt.Println("Depth: ", t1.Depth)
				} else if opc == 2 {
					fmt.Println("Depth: ", t2.Depth)
				}

			case 5:
				fmt.Println("In-Order")
				
				var opc int
				fmt.Println("1. T1 | 2.T2")
				fmt.Scanf("%d", &opc)
				if opc == 1{
					t1.InOrder()
				} else if opc == 2 {
					t2.InOrder()
				}

			case 6:
				fmt.Println("Pre-Order")
				
				var opc int
				fmt.Println("1. T1 | 2.T2")
				fmt.Scanf("%d", &opc)
				if opc == 1{
					t1.PreOrder()
				} else if opc == 2 {
					t2.PreOrder()
				}

			case 7:
				fmt.Println("Join Trees")	
				var opc int
				fmt.Println("1. Join 1 to 2 | 2. Join 2 to 1")
				fmt.Scanf("%d", &opc)
				if opc == 1 {
					t1.AddTree(t2)
				} else if opc == 2{
					t2.AddTree(t1)
				}
			
			default:
				
			
		}

	}
	
}
