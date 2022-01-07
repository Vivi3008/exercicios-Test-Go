package main

import (
	"fmt"
	"os"

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

	//ex extra 1
	ps := interfaces.Person{
		Name:   "Eu",
		Age:    35,
		Gender: "Feminino",
	}

	myFileOs, _ := os.OpenFile("../teste.json", os.O_RDONLY, os.ModeDevice)

	myFile := interfaces.DataJson{
		Data: *myFileOs,
	}

	interfaces.PrintInfo(ps)
	interfaces.PrintInfo(myFile)

	//ex extra 2
	resInt, err := interfaces.PrintType(int32(45))
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resInt)
	}
	resFloat, err := interfaces.PrintType(float32(45.1))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resFloat)
	}
	resStr, err := interfaces.PrintType("oi")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resStr)
	}
}
