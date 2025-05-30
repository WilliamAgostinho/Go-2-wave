package main

import (
	"fmt"
	"sort"
)

func main() {
	var numeros [12]int

	for i := 0; i < 12; i++ {
		fmt.Printf("Informe o %dº número: ", i+1)
		fmt.Scanln(&numeros[i])
	}

	slice := numeros[:]

	sort.Sort(sort.Reverse(sort.IntSlice(slice)))

	fmt.Println("Números em ordem decrescente:")
	for _, num := range slice {
		fmt.Println(num)
	}
}
