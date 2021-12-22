package funcoes

import "time"

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
