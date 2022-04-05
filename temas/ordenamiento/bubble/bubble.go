package bubble

import "fmt"

type dato struct {
	valor int
}

type contenedor struct {
	cont dato
}

func bubbleSort(c []contenedor) []contenedor{
	i := 0
	ordered := false
	for !ordered {
		ordered = true
		for i = 0; i<len(c) - 1; i++ {
			if c[i].cont.valor > c[i+1].cont.valor {
				aux := c[i]
				c[i] = c[i+1]
				c[i+1] = aux
				ordered = false
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

	a := bubbleSort(c);
	fmt.Println(c)
	fmt.Println(a)
	
}

