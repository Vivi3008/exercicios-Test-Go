package main

import (
	"fmt"
	"time"

	funcoes "github.com/Vivi3008/exercicios-Test-Go/Funcoes"
)

func main() {
	fmt.Printf("Ex 1 %s\n", funcoes.SayHello())
	fmt.Printf("Ex 1 %s\n", funcoes.SayHelloName("Lara Croft"))
	fmt.Printf("Ex 1 %s\n", funcoes.SayHelloPeriod("Jenny", time.Now()))
}
