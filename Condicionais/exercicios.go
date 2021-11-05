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
