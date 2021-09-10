package tiposcompostos

import "testing"

func TestExercicios(t *testing.T) {
	t.Run("Ex 1 - Deve verificar se é um array com 7 posições", func(t *testing.T) {
		result := criarArray()
		expected := [7]string{"Zero", "Um", "Dois", "Trẽs", "Quatro", "Cinco", "Seis"}

		if result != expected {
			t.Errorf("Resultado %v, esperado %v", result, expected)
		}
	})
}
