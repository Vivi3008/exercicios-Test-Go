package funcoes

import (
	"strings"
	"time"
)

//ex 1
func SayHello() string {
	return "Olar"
}

//ex 2
func SayHelloName(name string) string {
	return "Olar " + name
}

//ex 3
func SayHelloPeriod(name string, periodo time.Time) string {
	var period string
	hora, _, _ := periodo.UTC().Clock()

	switch {
	case hora > 0 && hora < 6:
		period = "Boa Madrugada"
	case hora >= 6 && hora < 12:
		period = "Bom Dia"
	case hora >= 12 && hora < 18:
		period = "Boa Tarde"
	case hora >= 18 && hora < 24:
		period = "Boa Noite"
	default:
		period = "Olar"
	}

	return period + " " + name
}

//ex 4 são todas as anteriores

//ex extra 1
func ModifyList(nums ...int) ([]int, int) {
	var list []int
	var total int

	for _, v := range nums {
		if v%2 != 0 {
			list = append(list, v*2)
		} else {
			list = append(list, v/2)
		}
	}

	for _, x := range list {
		total += x
	}

	return list, total
}

//ex extra 2
func CheckText(text string, letter string) int {
	count := 0
	for _, txt := range text {
		if string(txt) == letter || string(txt) == strings.ToUpper(letter) {
			count += 1
		}
	}
	return count
}

type Venda struct {
	Nome  string
	Valor int
	Data  time.Time
}

type ListVenda struct {
}

// ex extra 3
func ProcessSales(sales []Venda) map[string]int {
	var list = make(map[string]int, 0)

	for _, venda := range sales {
		valueAcc, ok := list[venda.Nome]
		if ok {
			list[venda.Nome] = valueAcc + venda.Valor
		} else {
			list[venda.Nome] = venda.Valor
		}
	}
	return list
}

// ex extra 4
func FormatWordForward(text string) string {
	var txt []byte
	sb := []byte(text)

	for _, v := range sb {
		switch {
		case v == 122: // caso as letras sejam x, y ou z
			txt = append(txt, 97+2)
		case v == 120:
			txt = append(txt, 97)
		case v == 121:
			txt = append(txt, 97+1)
		default:
			txt = append(txt, v+3)
		}
	}
	return string(txt)
}

func FormatWordBack(text string) string {
	var txt []byte
	sb := []byte(text)

	for _, v := range sb {
		switch {
		case v == 97: // caso as letras sejam a, b ou c
			txt = append(txt, 120)
		case v == 98:
			txt = append(txt, 121)
		case v == 99:
			txt = append(txt, 122)
		default:
			txt = append(txt, v-3)
		}
	}
	return string(txt)
}
