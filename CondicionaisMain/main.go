package main

import (
	"fmt"

	condicionais "github.com/Vivi3008/exercicios-Test-Go/Condicionais"
)

func main() {
	maior := condicionais.MaiorIdade(2013)

	if !maior {
		fmt.Println("Não Pode Votar!")
	} else {
		fmt.Println("Pode Votar!")
	}

	numero := condicionais.Positivo(-66)

	if !numero {
		fmt.Println("Número negativo")
	} else {
		fmt.Println("Número positivo")
	}
}
