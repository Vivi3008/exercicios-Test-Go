package tiposdados

import "testing"

func TestExercicios(t *testing.T) {

	t.Run("Ex 1: Retorna um valor de variavel e seu tipo", func(t *testing.T) {
		result := PrintNum(2)
		expected := 2
		resultName := ImprimeNome("Vivi")
		expectedName := "Olá Vivi"
		floatResult := PrintFloat(45.5)
		expectedFloat := 45.5

		if result != expected {
			t.Errorf("Resultado %v, tipo %T esperado %v!\n", result, result, expected)
		}
		if resultName != expectedName {
			t.Errorf("Resultado %s, tipo %T, esperado %s!\n", resultName, resultName, expectedName)
		}
		if floatResult != float32(expectedFloat) {
			t.Errorf("Resultado %.2f, tipo %T, esperado %.2f!\n", floatResult, floatResult, expectedFloat)
		}
	})

	t.Run("Ex 2 Somar dois numeros", func(t *testing.T) {
		a := 230
		b := 27
		total := Sum(a, b)
		expected := 257

		if total != expected {
			t.Errorf("Resultado %d, valor esperado %d\n", total, expected)
		}
	})

	t.Run("Ex 3 Lista de compras", func(t *testing.T) {
		listaPrecos := map[string]float64{
			"Banana":     1.25,
			"Cerveja":    2.98,
			"Abacate":    4.59,
			"Salgadinho": 7.29,
		}
		total := CalculaCompras(listaPrecos)
		expected := 68.396

		if total != float64(expected) {
			t.Errorf("Resultado %.4f, valor esperado %.4f\n", total, expected)
		}
	})

}
