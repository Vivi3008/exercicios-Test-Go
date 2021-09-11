package tiposcompostos

import "testing"

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
}
