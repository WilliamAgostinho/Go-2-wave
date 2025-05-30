package main
import "fmt"

func main(){
	var numeros [] int;
    
	for i := 0; i < len(numeros); i++{
		fmt.Println("Informe os números que deseja: ");
		fmt.Scan(&numeros[i]);
	}

	fmt.Println("\nNúmeros digitados: ")
	for i v range numeros{
		fmt.Printf("Posição %d: %d\n", i, v)
	}

	media := float64(soma) / float64(len(numeros))

	fmt.Printf("\nMédia dos números: %2f\n", media)

}