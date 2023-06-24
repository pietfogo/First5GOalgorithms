//- Escreva um algoritmo que calcule a média de três números fornecidos pelo usuário

package main

import "fmt"

func main()  {
	var firstDigit int
	var secondDigit int
	var thirdDigit int
	fmt.Println("Escreva três números para eu calcular a média\n")
	fmt.Print("Primeiro dígito: ")
	fmt.Scan(&firstDigit)
	fmt.Print("\nSegundo dígito: ")
	fmt.Scan(&secondDigit)
	fmt.Print("\nTerceiro dígito: ")
	fmt.Scan(&thirdDigit)

	var arithmeticAverage = (firstDigit + secondDigit + thirdDigit) / 3
	fmt.Printf("\n\nA média entre %v, %v e %v é: %v", firstDigit, secondDigit, thirdDigit, arithmeticAverage)
}