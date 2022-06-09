package insertion

import "fmt"

type dato struct {
	valor int
}

type contenedor struct {
	cont dato
}

func insertionSort(c []contenedor)[]contenedor{
	var pivote contenedor
	var j int
	for i := 0; i< len(c); i++ {
		pivote = c[i]
		j = i-1
		for j>=0 && c[j].cont.valor > pivote.cont.valor {
			c[j+1] = c[j]
			j--
		}
		c[j+1] = pivote
	}
	return c
}

func main() {
	c := []contenedor{
		{dato{10}},
		{dato{2}},
		{dato{9}},
		{dato{15}},
		{dato{8}},
	}

	c = insertionSort(c)
	fmt.Println(c)
	
}

