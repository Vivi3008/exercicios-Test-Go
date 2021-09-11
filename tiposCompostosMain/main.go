package main

import (
	"fmt"

	tiposcompostos "github.com/Vivi3008/exercicios-Test-Go/tiposCompostos"
)

func main() {
	//ex 1 Imprime um array
	fmt.Printf("O tipo do array é %T\n", tiposcompostos.CriarArray())

	//ex 2 Atribui um array a outro
	fmt.Printf("O array é %v\n", tiposcompostos.SetTwoArrays())

	//ex 3 Imprime os tres arrays
	fmt.Printf("Arrays %v\n", tiposcompostos.AddTwoSlices())
}
