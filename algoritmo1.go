package main

import "fmt"

func main()  {
	var parint int
	fmt.Println("Digite um número e eu direi se é par ou ímpar: ")
	fmt.Scan(&parint)

	if parint % 2 == 0{
		fmt.Printf("O %v é PAR", parint)
	} else {
		fmt.Printf("O %v é ÍMPAR", parint)
	}
	
}