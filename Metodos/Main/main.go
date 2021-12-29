package main

import (
	"fmt"

	met "github.com/Vivi3008/exercicios-Test-Go/Metodos"
)

func main() {
	//ex 1
	var y = met.Circle{
		Raio: 32.1,
	}

	fmt.Printf("A area do circulo é %.2f\n", y.CalcArea())
	fmt.Printf("O perimetro do circulo é %.2f\n", y.CalcPerimetro())

	//ex 2
	list := met.ListNumbers{4, 6, 8, 9, 4, 5, 7, 10}
	fmt.Printf("Soma da lista %v\n", list.CalcSoma())
	fmt.Printf("Media da lista %.2f\n", list.CalcMedia())

	//ex extra 1
	firsPilha := met.Pilha{1, 2, 3, 4}
	result := met.Pilha(firsPilha.Push(5))
	fmt.Printf("Adicionando um item a pilha %v\n", result)
	fmt.Printf("Removendo o ultimo item %v\n", result.Pop())
}
