package main

import (
	"fmt"
	"time"

	funcoes "github.com/Vivi3008/exercicios-Test-Go/Funcoes"
)

func main() {
	fmt.Printf("Ex 1 %s\n", funcoes.SayHello())
	fmt.Printf("Ex 2 %s\n", funcoes.SayHelloName("Lara Croft"))
	fmt.Printf("Ex 3 %s\n", funcoes.SayHelloPeriod("Jenny", time.Now()))
	x, y := funcoes.ModifyList(1, 6, 5, 8, 6, 7)
	fmt.Printf("Ex 1 Extra Lista Modificada: %v Soma lista: %v\n", x, y)
	fmt.Printf("Ex 2 Extra Quantidade da letra a: %v\n", funcoes.CheckText("Eu amo Golaaaang", "a"))
	vendas := []funcoes.Venda{
		{
			Nome:  "Tete",
			Valor: 653,
			Data:  time.Date(2019, 6, 5, 11, 51, 04, 0, time.UTC),
		},
		{
			Nome:  "TiTI",
			Valor: 658,
			Data:  time.Date(2019, 6, 8, 21, 51, 04, 0, time.UTC),
		},
		{
			Nome:  "Eita",
			Valor: 100,
			Data:  time.Date(2019, 6, 8, 21, 51, 04, 0, time.UTC),
		},
		{
			Nome:  "Eita",
			Valor: 200,
			Data:  time.Date(2019, 6, 8, 21, 51, 04, 0, time.UTC),
		},
		{
			Nome:  "Gigi",
			Valor: 741,
			Data:  time.Now(),
		},
		{
			Nome:  "Gigi",
			Valor: 100,
			Data:  time.Now(),
		},
	}
	fmt.Printf("Lista de vendas %v\n", funcoes.ProcessSales(vendas))

	name := "viviane"
	name = funcoes.FormatWordForward(name)
	fmt.Printf("Altera cada letra da palavra para 3 caracteres seguintes: %s\n", name)
	name = funcoes.FormatWordBack(name)
	fmt.Printf("Conserta a palavra: %s\n", name)
}
