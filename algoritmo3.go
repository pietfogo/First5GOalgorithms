package main

import "fmt"

func main()  {
	var anoBissexto int
	var addictionalVal = 2
	fmt.Print("Digite o ano que você está: ")
	fmt.Scan(&anoBissexto)
	anoBissexto = anoBissexto + addictionalVal
	if anoBissexto % 4 == 0 {
		fmt.Printf("\nO ano é bissexto e vai ter ou já teve copa hehehe")
	} else {
		fmt.Printf("\n O ano não é bissexto")
	}
}