package tiposcompostos

import (
	"fmt"
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
			"Azul":    "EKJ6655",
			"Laranja": "KJI6655",
			"Marelo":  "54dfa54",
		}

		ExcluirValorMapa(cores, "Azul")

		_, exist := cores["Azul"]

		if exist {
			t.Errorf("A cor %v nao foi excluida", exist)
		}
	})

	t.Run("Ex 6 Mapa de meses do ano", func(t *testing.T) {
		mesesAno := map[int]string{
			1:  "Janeiro",
			2:  "Fevereiro",
			3:  "Março",
			4:  "Abril",
			5:  "Maio",
			6:  "Junho",
			7:  "Julho",
			8:  "Agosto",
			9:  "Setembro",
			10: "Outubro",
			11: "Novembro",
			12: "Dezembro",
		}
		expected := 12

		if len(mesesAno) != expected {
			t.Errorf("Resultado %v, esperado %v", len(mesesAno), expected)
		}
	})

	t.Run("Ex extra 1, Array de times", func(t *testing.T) {
		timeAmarelo := []string{"Fernando", "João", "Lúcia", "Mariana", "Ana"}
		timeVermelho := []string{"Helena", "Jonas", "José", "Juliana"}

		resultTimeAmarelo, err := ImprimeJogadores(timeAmarelo)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		resultTimeVermelho, err := ImprimeJogadores(timeVermelho)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		if resultTimeAmarelo[0] != "Fernando" {
			t.Errorf("Expected %s, got %s", "Fernando", resultTimeAmarelo[0])
		}

		if resultTimeVermelho[0] != "Helena" {
			t.Errorf("Expected %s, got %s", "Helena", resultTimeVermelho[0])
		}

		// ex extra 2
		timeVermelho = append(timeVermelho, "Luis Inácio")

		resultTimeVermelho, err = ImprimeJogadores(timeVermelho)

		if err != nil {
			t.Errorf("Expected nil, got %s", err.Error())
		}

		lastPosition := len(timeVermelho) - 1

		if resultTimeVermelho[lastPosition] != "Luis Inácio" {
			t.Errorf("Expected %s, got %s", "Luis Inácio", resultTimeVermelho[lastPosition])
		}
	})

	t.Run("Ex Extra 3", func(t *testing.T) {
		countrys := map[string]string{
			"Goiânia":   "Brasil",
			"São Paulo": "Brasil",
			"Tokio":     "Japan",
			"Lisboa":    "Portugal",
			"New York":  "US",
			"Nairobi":   "Quenia",
		}

		country := "Brasil"

		appear := AppearCountry(countrys, country)

		fmt.Printf("Appear %v", appear)

		if appear[country] != 2 {
			t.Errorf("Expected %v, got %v", 2, appear[country])
		}
	})
}
