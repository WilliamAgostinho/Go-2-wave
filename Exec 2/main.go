package main

import (
	"fmt"
)

func main() {
	var med1, med2, med3, med4 float64
	var media float64

	println("Informe a primeira média: ")
	fmt.Scanln(&med1)

	println("Informe a segunda média: ")
	fmt.Scanln(&med2)

	println("Informe a terceira média: ")
	fmt.Scanln(&med3)

	println("Informe a quarta média: ")
	fmt.Scanln(&med4)

	media = (med1 + med2 + med3 + med4) / 4

	fmt.Printf("Média: %.2f\n", media)

	if media >= 5 {
		print("Aluno aprovado")
	} else {
		print("Aluno reprovado")
	}
}
