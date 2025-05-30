package main

import "fmt"

func main() {
	var a, b, c, resultado int

	fmt.Print("Informe o primeiro valor: ")
	fmt.Scanln(&a)

	fmt.Print("Informe o segundo valor: ")
	fmt.Scanln(&b)

	fmt.Print("Informe o terceiro valor: ")
	fmt.Scanln(&c)

	resultado = a + b + c
	resultado = resultado * resultado

	fmt.Println("Resultado:", resultado)
}
