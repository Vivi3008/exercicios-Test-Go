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
	//ex 3 e 4
	fmt.Printf("Faixa de idade: %s\n", condicionais.VerifyAge(18))
	fmt.Printf("Faixa de idade: %s\n", condicionais.VerifyAgeSwitch(85))

	//ex 5
	fmt.Printf("Periodo do dia: %s\n", condicionais.VerifyDayPeriod(4))

	//ex extra1
	fmt.Printf("Maior numero: %v\n", condicionais.CheckBigger(1, 4, 3))
	//ex extra2
	fmt.Println(condicionais.CheckMultiple(0))
}
