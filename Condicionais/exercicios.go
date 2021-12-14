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
	} else {
		return "Idoso"
	}
}
