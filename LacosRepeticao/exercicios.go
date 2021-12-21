package lacosRepeticao

import (
	"bufio"
	"fmt"
	"os"
)

type Ape struct {
	Num  int
	Nome string
	Vaga int
}

var apartamentos = []Ape{
	{
		Num:  1,
		Nome: "Giselle",
		Vaga: 2,
	},
	{
		Num:  2,
		Nome: "Vanny",
		Vaga: 1,
	},
	{
		Num:  3,
		Nome: "Leticia",
		Vaga: 2,
	},
}

var txtResposta = "Agora sim podemos dividir igualmente entre mim e você!"

func NumberSequence(begin int, end int) []int {
	numbers := make([]int, 0)

	for i := begin; i <= end; i++ {
		numbers = append(numbers, i)
	}

	return numbers
}

func DayHours() ([]int, []int) {
	hours := make([]int, 0)
	minutes := make([]int, 0)

	for i := 0; i < 24; {
		hours = append(hours, i)
		for m := 0; m < 60; m++ {
			minutes = append(minutes, m)
		}
		i++
	}

	return hours, minutes
}

func SliceString() []string {
	strings := []string{"um", "dois", "tres", "quatro", "cinco"}

	return strings
}

func ShopList() []string {
	list := []string{"Arroz", "Carne", "Chocolate", "Chicletes", "Verduras", "Refrigerante"}

	return list
}

func PrintValues() {
	arr := []int{}
	for i := 1; i <= 10; i++ {
		arr = append(arr, i)
		fmt.Printf("%v\n", arr)
	}
}

func PrintApes() {
	for _, ap := range apartamentos {
		fmt.Printf("Num Ap: %v | Nome: %s | Vaga: %v\n", ap.Num, ap.Nome, ap.Vaga)
	}
}

func LerNumeroPar() {
	num := 1
	for num%2 != 0 {
		fmt.Println("Digite um numero:")
		fmt.Fscanf(os.Stdin, "%d", &num)
	}
	fmt.Println(txtResposta)
}

func LerTexto() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}

var letras = map[string]int{
	"a": 0,
	"b": 0,
	"c": 0,
	"d": 0,
	"e": 0,
	"f": 0,
	"g": 0,
	"h": 0,
	"i": 0,
	"j": 0,
	"k": 0,
	"l": 0,
	"m": 0,
	"n": 0,
	"p": 0,
	"q": 0,
	"r": 0,
	"s": 0,
	"t": 0,
	"u": 0,
	"v": 0,
	"w": 0,
	"x": 0,
	"y": 0,
	"z": 0,
}

func VerifyLetters() {
	text := LerTexto()

	for _, v := range text {
		switch string(v) {
		case "a":
			letras["a"] += 1
		case "b":
			letras["b"] += 1
		case "c":
			letras["c"] += 1
		case "d":
			letras["d"] += 1
		case "e":
			letras["e"] += 1
		case "f":
			letras["f"] += 1
		case "g":
			letras["g"] += 1
		case "h":
			letras["h"] += 1
		case "i":
			letras["i"] += 1
		case "j":
			letras["j"] += 1
		case "k":
			letras["k"] += 1
		case "l":
			letras["l"] += 1
		case "m":
			letras["m"] += 1
		case "n":
			letras["n"] += 1
		case "p":
			letras["p"] += 1
		case "q":
			letras["q"] += 1
		case "r":
			letras["s"] += 1
		case "t":
			letras["u"] += 1
		case "v":
			letras["v"] += 1
		case "w":
			letras["w"] += 1
		case "x":
			letras["x"] += 1
		case "y":
			letras["y"] += 1
		case "z":
			letras["z"] += 1
		}
	}
	for i, v := range letras {
		fmt.Printf("%s => %d ", i, v)
	}
}
