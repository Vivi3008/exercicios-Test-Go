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

		if total != expected {
			t.Errorf("Resultado %.4f, valor esperado %.4f\n", total, expected)
		}
	})

	t.Run("Ex 4 Retornar uma frase com nome e cor favorita", func(t *testing.T) {
		result := CorFavorita("Viviane", "Roxo")
		expected := "Meu nome é Viviane minha cor favorita é Roxo"

		if result != expected {
			t.Errorf("Resultado %s, esperado %s", result, expected)
		}
	})

	t.Run("Ex 5 Operaçoes relacionais", func(t *testing.T) {
		a := 35
		b := 14
		result := EMaior(a, b)
		expected := true

		if result != expected {
			t.Errorf("Resultado %t, esperado %t", result, expected)
		}
	})

	t.Run("Ex 6 Operador &&", func(t *testing.T) {
		a := 1
		b := 4
		c := 3
		result := VerificaMaior(a, b, c)
		expected := "B é maior"

		if result != expected {
			t.Errorf("Resultado %s, esperado %s", result, expected)
		}
	})

	t.Run("Ex. Extra 1", func(t *testing.T) {
		result := VerifyKm(8)
		expected := 8

		if result != int32(expected) {
			t.Errorf("Resultado %v, esperado %v", result, expected)
		}
	})

	t.Run("Ex Extra 2", func(t *testing.T) {
		result := CalculaAniversario(1986)
		expected := 35

		if result != expected {
			t.Errorf("Resultado %d, esperado %d", result, expected)
		}
	})

}
