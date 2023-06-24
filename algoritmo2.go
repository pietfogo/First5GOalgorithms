package main

import "fmt"

func main()  {
	var readNumber int
	fmt.Print("Digite um número: ")
	fmt.Scan(&readNumber)

	if readNumber > 0 {
		fmt.Printf("É positivo")
	} else if readNumber < 0 {
		fmt.Printf("É negativo")
	} else {
		fmt.Printf("O número é igual a %v", readNumber)
	}
	
}