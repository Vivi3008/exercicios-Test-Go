package tiposcompostos

import (
	"testing"
)

func TestExercicios(t *testing.T) {
	t.Run("Ex 1 - Deve verificar se é um array com 7 posições", func(t *testing.T) {
		result := CriarArray()
		expected := 7

		if len(result) != expected {
			t.Errorf("Resultado %v, esperado %v", result, expected)
		}
	})

	t.Run("Ex 2 - Verifica se um array foi atribuido", func(t *testing.T) {
		result := SetTwoArrays()
		expected := [4]string{"Dourado", "Roxo", "Rosa", "Azul"}

		if result != expected {
			t.Errorf("Resultado %v, esperado %v", result, expected)
		}
	})

	t.Run("Ex 3 - Verifica tamanho do array", func(t *testing.T) {
		result := AddTwoSlices()
		expected := 12

		if len(result) != expected {
			t.Errorf("Resultado %v, experado %v", len(result), expected)
		}
	})

	t.Run("Ex 4 - Deve retornar um slice de string", func(t *testing.T) {
		result := SliceLiteral()
		sizeExpected := 8

		if len(result) != sizeExpected {
			t.Errorf("Resultado %v, esperado %v", len(result), sizeExpected)
		}
	})

	t.Run("Ex 5 Excluir um valor de um mapa e verificar se existe", func(t *testing.T) {
		cores := map[string]string{
			"Azul" : "EKJ6655",
			"Laranja" : "KJI6655",
			"Marelo" : "54dfa54",
		}

		ExcluirValorMapa(cores, "Azul")

		_, exist := cores["Azul"]

		if exist {
			t.Errorf("A cor %v nao foi excluida", exist)
		}
	})

	t.Run("Ex 6 Mapa de meses do ano", func(t *testing.T) {
		mesesAno := map[int]string{
			1 : "Janeiro",
			2 : "Fevereiro",
			3 : "Março",
			4 : "Abril",
			5 : "Maio",
			6 : "Junho",
			7: "Julho",
			8 : "Agosto",
			9: "Setembro",
			10 : "Outubro",
			11: "Novembro",
			12 : "Dezembro",
		}
		expected := 12

		if len(mesesAno) != expected {
			t.Errorf("Resultado %v, esperado %v", len(mesesAno), expected)
		}
	})
}
