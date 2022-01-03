package metodos

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

// ex 1
type Circle struct {
	Raio float64
}

func (c Circle) CalcArea() float64 {
	area := math.Pi * math.Pow(c.Raio, 2)

	return area
}

func (c Circle) CalcPerimetro() float64 {
	perim := 2 * (math.Pi * c.Raio)

	return perim
}

//ex 2
type ListNumbers []int

func (l ListNumbers) CalcSoma() int {
	var count int
	for _, v := range l {
		count += v
	}
	return count
}

func (l ListNumbers) CalcMedia() float64 {
	media := float64(l.CalcSoma()) / float64(len(l))

	return media
}

//ex extra 1
type Pilha []int

func (p Pilha) Push(value int) []int {
	p = append(p, value)

	return p
}

func (p Pilha) Pop() []int {
	var res []int
	for i, v := range p {
		if len(p)-1 != i {
			res = append(res, v)
		}
	}
	return res
}

//ex extra 2
type CPF string

var (
	ErrInvalidCpf       = errors.New("invalid Cpf")
	ErrInvalidCaracters = errors.New("cpf has less than 11 caracters or have letters")
)

func (c CPF) ValidaCpf() ([]int, error) {
	if len(c) != 11 || c.haveLetters() {
		return []int{}, ErrInvalidCaracters
	}

	latTwoDigits := joinStringsToInt(string(c[9]), string(c[10]))

	nineDigits := cpfToNineSliceNumbers(c)
	peso1 := []int{10, 9, 8, 7, 6, 5, 4, 3, 2}

	firstDigit := calcNextDigit(peso1, nineDigits)

	tenDigits := append(nineDigits, firstDigit)
	peso2 := []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2}

	secondDigit := calcNextDigit(tenDigits, peso2)

	cpfComplete := append(tenDigits, secondDigit)

	lastDigits := joinStringsToInt(fmt.Sprint(firstDigit), fmt.Sprint(secondDigit))

	if lastDigits != latTwoDigits {
		return []int{}, ErrInvalidCpf
	}

	return cpfComplete, nil
}

func (c CPF) haveLetters() bool {
	var letters bool
	for _, v := range c {
		if v < 48 || v > 57 {
			letters = true
		}
	}
	return letters
}

func cpfToNineSliceNumbers(cpf CPF) []int {
	var cpfNumbers []int

	for i := 0; i < 9; i++ {
		num, _ := strconv.ParseInt(string(cpf[i]), 10, 64)
		cpfNumbers = append(cpfNumbers, int(num))
	}

	return cpfNumbers
}

func joinStringsToInt(x string, y string) int {
	strJoin := fmt.Sprintf("%v%v", x, y)
	latTwoDigits, _ := strconv.ParseInt(strJoin, 10, 64)
	return int(latTwoDigits)
}

func calcNextDigit(slice1 []int, slice2 []int) int {
	var sum int
	var digit int
	for i, v := range slice1 {
		mult := (v * slice2[i])
		sum += mult
	}

	rest := sum % 11

	if rest < 2 {
		digit = 0
	} else {
		digit = 11 - rest
	}
	return digit
}
