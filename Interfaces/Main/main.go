package main

import (
	"fmt"

	interfaces "github.com/Vivi3008/exercicios-Test-Go/Interfaces"
)

func main() {
	//ex 1
	circ := interfaces.Circulo{
		Raio: 5,
	}
	quad := interfaces.Quadrado{
		Lado: 9,
	}
	fmt.Printf("A area do circulo é %.2f\n", interfaces.CalculaAreaGeometrica(circ))
	fmt.Printf("A area do quadrado é %.2f\n", interfaces.CalculaAreaGeometrica(quad))

	//ex 2
	dog := interfaces.Cachorro{
		Tamanho:   "Pequeno",
		Som:       "AU AU",
		Brinquedo: "Franguinho",
		Comida:    "Ração",
	}
	cat := interfaces.Gato{
		Som:       "Miauuuuu",
		Brinquedo: "Bolinha",
		Comida:    "Sachê",
	}

	fmt.Println(interfaces.ApresentaAnimal(cat))
	fmt.Println(interfaces.ApresentaAnimal(dog))
}
