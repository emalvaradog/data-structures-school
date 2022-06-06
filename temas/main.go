package main

import (
	"fmt"
	"topics/estructuras/tree"
)

func main() {

	t := tree.Tree{}

	var opc int
	for opc != 10 {
		fmt.Println("Menu")
		fmt.Println("1. AddNode")
		fmt.Println("2. SearchNode")
		fmt.Println("3. DeleteNode")
		fmt.Println("4. Depth")
		fmt.Println("5. In-Order")
		fmt.Println("6. Pre-Order")
		fmt.Println("7. Post-Order")
		fmt.Println("8. MinNode")
		fmt.Println("9. MaxNode")
		fmt.Println("10. Leave")
		fmt.Scanf("%d", &opc)

		switch opc {
			case 1:
				data := tree.TData{}
				fmt.Println("Enter int value")
				fmt.Scanf("%d", &data.Value)
				t.AddNode(data)
				fmt.Println("Node Added")

			case 2:
				data := tree.TData{}
				fmt.Println("Enter int value")
				fmt.Scanf("%d", &data.Value)
				res := t.SearchNode(data)
				if res == nil {
					fmt.Println("Node not found")
				}

				fmt.Println("Node found")
				fmt.Printf("Value: %d | Left: %d | Right: %d \n", res.Data.Value, res.Left.Data.Value, res.Right.Data.Value)

			case 3:	
				data := tree.TData{}
				fmt.Println("Enter int value")
				fmt.Scanf("%d", &data.Value)
				t.DeleteNode(data)
				fmt.Println("Node deleted")

			case 4:
				fmt.Println("Depth: ", t.Depth())

			case 5:
				fmt.Println("In-Order")
				t.InOrder()

			case 6:
				fmt.Println("Pre-Order")
				t.PreOrder()

			case 7:
				fmt.Println("Post-Order")
				t.PostOrder()

			case 8:
				fmt.Println("Min Node")
				res := t.MinNode()
				fmt.Println(res)

			case 9:
				fmt.Println("Max Node")
				res := t.MaxNode()
				fmt.Println(res)

			default:
				
			
		}

	}
	
}
