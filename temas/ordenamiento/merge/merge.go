package merge

import "fmt"

type dato struct {
	valor int
}

type contenedor struct {
	cont dato
}

func mergeSort(c []contenedor) []contenedor {
	if len(c) < 2 {
		return c
	}
	mid := (len(c) / 2)
	c = merge(mergeSort(c[:mid]), mergeSort(c[mid:]))
	return c
}

func merge(a, b []contenedor) []contenedor {
	size, i, j := len(a) + len(b), 0, 0
	final := make([]contenedor, size)

	for k := 0; k < size; k++ {
		if i > len(a) - 1 && j <= len(b) - 1 {
			final[k] = b[j]
			j++
		} else if j > len(b) -1 && i <= len(a)-1{
			final[k] = a[i]
			i++
		} else if a[i].cont.valor < b[j].cont.valor {
			final[k] = a[i]
			i ++
		} else {
			final[k] = b[j]
			j++
		}
	}
	return final
}

func main() {
	c := []contenedor{
		{dato{10}},
		{dato{2}},
		{dato{9}},
		{dato{15}},
		{dato{8}},
		{dato{1}},
	}

	c = mergeSort(c)
	fmt.Println(c)
	
}


