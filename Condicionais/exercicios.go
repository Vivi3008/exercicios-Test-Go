package condicionais

import "time"

var dataAtual = time.Now()

func MaiorIdade(nascimento int) bool {
	idade := dataAtual.Year() - nascimento

	return idade >= 18
}

func Positivo(num int) bool {
	return num > 0
}

func VerifyAge(age int) string {
	if age < 18 {
		return "Menor de idade"
	} else if age >= 18 && age <= 60 {
		return "Maior de idade"
	} else if age < 0 {
		return "Idade não pode ser menor que zero"
	} else {
		return "Idoso"
	}
}

func VerifyAgeSwitch(age int) string {
	rangeAge := ""

	switch {
	case age < 0:
		rangeAge = "Idade não pode ser menor que zero"
	case age < 18:
		rangeAge = "Menor de idade"
	case age >= 18 && age <= 60:
		rangeAge = "Maior de idade"
	case age > 60:
		rangeAge = "Idoso"
	}

	return rangeAge
}

func VerifyDayPeriod(hour int) string {
	period := ""

	switch {
	case hour > 0 && hour < 6:
		period = "Madrugada"
	case hour >= 6 && hour < 12:
		period = "Manhã"
	case hour >= 12 && hour < 18:
		period = "Tarde"
	case hour >= 18 && hour < 24:
		period = "Noite"
	default:
		period = "Valor informado inválido"
	}

	return period
}

func CheckBigger(a int, b int, c int) int {
	var bigger int

	if a > b && a > c {
		bigger = a
	} else if b > c {
		bigger = b
	} else {
		bigger = c
	}

	return bigger
}

func CheckMultiple(num int) string {
	var value string

	if num == 0 {
		return "Zero"
	}

	switch {
	case num%2 == 0:
		value = "Multiplo de 2"
	case num%3 == 0:
		value = "Multiplo de 3"
	default:
		value = "Nenhuma das opções"
	}

	return value
}
