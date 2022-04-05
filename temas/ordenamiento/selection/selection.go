package selection

import "fmt"

type dato struct {
	valor int
}

type contenedor struct {
	cont dato
}

func selectionSort(c []contenedor) []contenedor{
	for i:= 0; i < len(c); i++ {
		for j:= i+1; j < len(c); j++ {
			if c[j].cont.valor < c[i].cont.valor {
				cont := c[i]
				c[i] = c[j]
				c[j] = cont
			}
		}
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

	selectionSort(c)
	fmt.Println(c)
	
}

