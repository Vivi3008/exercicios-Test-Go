package main

import (
	"fmt"

	tiposdados "github.com/Vivi3008/exercicios-Test-Go/TiposDados"
)

func main() {
	//ex 1
	fmt.Printf("O valor da variavel é %d\n", tiposdados.PrintNum(6))
	fmt.Printf("%s\n", tiposdados.ImprimeNome("Viviane"))
	fmt.Printf("Imprime um float com duas casas decimais %.2f\n", tiposdados.PrintFloat(56.45556))

	//ex 2
	fmt.Printf("A soma dos valores 230 e 27 é %v\n", tiposdados.Sum(230, 27))

	//ex 3
	listaCompras := map[string]float64{
		"Banana":     65.6,
		"Cerveja":    85.2,
		"Abacate":    65.9,
		"Salgadinho": 1.22,
	}
	fmt.Printf("Valor total das compras: %.2f\n", tiposdados.CalculaCompras(listaCompras))

	//ex 4
	fmt.Printf("%s\n", tiposdados.CorFavorita("Viviane", "Roxo"))

	//ex 5
	fmt.Printf("65 é maior que 8? %t\n", tiposdados.EMaior(65, 8))

	//ex 6
	fmt.Printf("Qual é o maior numero? A = 6, B = 8, C = 10? %v\n", tiposdados.VerificaMaior(6, 8, 10))

	//ex extra 2
	idade := tiposdados.CalculaAniversario(1994)
	fmt.Printf("Sua idade é %v\n", idade)
}
