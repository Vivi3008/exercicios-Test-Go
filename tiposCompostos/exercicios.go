package tiposcompostos

type Arrays struct {
	arrayUm []int16
	arrayDois []int16
	arrayTres []int16
}

//funcao ex 1
func CriarArray() [7]string {

	x := [7]string{"Zero", "Um", "Dois", "Trẽs", "Quatro", "Cinco", "Seis"}
	arrayAleatorio := x

	return arrayAleatorio
}

//funcao ex 2
func SetTwoArrays() [4]string {
	x := [4]string{"Banana", "Maça", "Chocolate", "Cerveja"}
	y := [4]string{"Dourado", "Roxo", "Rosa", "Azul"}

	x = y

	return x
}

//ex 3
func AddTwoSlices() []int16 {
	sliceUm :=  []int16{1,2,3,4,5,6}
	sliceDois :=  []int16{7,8,9,10,11,12}

	var sliceTres []int16

	sliceTres = append(sliceUm, sliceDois...)

	return sliceTres
}

//ex 4
func SliceLiteral() []string {
	compras := []string{"Cerveja", "Chocolate", "Lasanha", "Bolacha", "Laranja"}

	compras = append(compras, "Cenoura", "Beterraba", "Maça")

	return compras
}

//ex 5
func ExcluirValorMapa(maps map[string]string, cor string){
	delete(maps, cor)
}

//ex 6
func MesesAno(meses map[int]string) map[int]string {
	primeiro := meses[1]
	ultimo := meses[12]

	result := map[int]string{
		1: primeiro,
		12 : ultimo,
	}

	return result
}

