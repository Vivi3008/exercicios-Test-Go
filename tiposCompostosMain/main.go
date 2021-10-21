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

	//ex 4 Imprime um slice literal
	fmt.Printf("Slice literal %v\n", tiposcompostos.SliceLiteral())

	//ex5
	cores := map[string]string{
		"Azul":    "EKJ6655",
		"Laranja": "KJI6655",
		"Marelo":  "54dfa54",
	}
	fmt.Printf("Mapa original %v\n", cores)

	tiposcompostos.ExcluirValorMapa(cores, "Azul")

	fmt.Printf("Mapa depois de excluir um valor %v\n", cores)

	//ex 6
	meses := map[int]string{
		1:  "Janeiro",
		2:  "Fevereiro",
		3:  "Março",
		4:  "Abril",
		5:  "Maio",
		6:  "Junho",
		7:  "Julho",
		8:  "Agosto",
		9:  "Setembro",
		10: "Outubro",
		11: "Novembro",
		12: "Dezembro",
	}
	fmt.Printf("Meses: %v\n", tiposcompostos.MesesAno(meses))

	//ex 7
	pessoa := tiposcompostos.Pessoa{Nome: "Viviane", Idade: 35}
	pessoa2 := tiposcompostos.Pessoa{Nome: "Giovanna", Idade: 9}
	pessoa3 := tiposcompostos.Pessoa{Nome: "Davy", Idade: 7}

	fmt.Printf("Pessoa 1: %v\n", tiposcompostos.ImprimePessoa(pessoa))
	fmt.Printf("Pessoa 2: %v\n", tiposcompostos.ImprimePessoa(pessoa2))
	fmt.Printf("Pessoa 3: %v\n", tiposcompostos.ImprimePessoa(pessoa3))

	timeAmarelo := []string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
	timeVermelho := []string{"Helena", "Jonas", "José", "Juliana"}

	timeTest, _ := tiposcompostos.ImprimeJogadores(timeAmarelo)
	timeTestVermelho, _ := tiposcompostos.ImprimeJogadores(timeVermelho)

	fmt.Printf("Time Amarelo: %v\n", timeTest)
	fmt.Printf("Time Vermelho: %v\n", timeTestVermelho)
}
