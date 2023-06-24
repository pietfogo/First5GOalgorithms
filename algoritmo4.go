package main 

import ("fmt"
		"strings")

func main()  {
	//var recebeVal string
	//fmt.Println("Escreva uma palavra e eu direi se é ou não um palíndromo: ")
	stringVal := "ovo"
	convertString := strings.Split(stringVal, "")
	numberElements := len(convertString)
	var arrayVal [] string
	arrayVal = make([]string, numberElements)
	//fmt.Printf("Número de elementos: %v, array vazio: %v, array cheio %v", numberElements, array, convertString)
	x := -1
	for i := numberElements-1; i != -1; i-- {
		x++
		arrayVal[x] = convertString[i]
	}
	//fmt.Printf("arrayVal: %v, convertString: %v", arrayVal, convertString)
	var cond bool
	for c := 0; c < numberElements; c++ {
		if arrayVal[c] != convertString[c]{
			cond = false
		} else {
			cond = true
		}
	}
	if cond == false {
		fmt.Println("Não é um palíndromo")
	} else{
		fmt.Println("É um palíndromo")
	}
}