package tiposdados

//funcoes exercicio 1
func PrintNum(num int) int {
	return num
}

func ImprimeNome(name string) string {
	return "Olá " + name
}

func PrintFloat(num float32) float32 {
	return num
}

//funcao ex 2
func Sum(a int, b int) int {
	return a + b
}

//funcao ex 3
func CalculaCompras(lista map[string]float64) float64 {
	banana := lista["Banana"] * 2.170
	cerveja := lista["Cerveja"] * 6
	abacate := lista["Abacate"] * 5.650
	salgadinho := lista["Salgadinho"] * 3

	total := banana + cerveja + abacate + salgadinho
	return float64(total)
}
