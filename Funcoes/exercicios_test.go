package funcoes

import (
	"testing"
	"time"
)

func TestExerciciosFuncoes(t *testing.T) {
	t.Run("Ex 1 Printa um cumprimento", func(t *testing.T) {
		cumprimento := SayHello()

		if cumprimento != "Olar" {
			t.Errorf("Expected Olar, got %s", cumprimento)
		}
	})

	t.Run("Ex 2 Printar nome no cumprimento", func(t *testing.T) {
		cumprimento := SayHelloName("Viviane")

		if cumprimento != "Olar Viviane" {
			t.Errorf("Expected Olar, got %s", cumprimento)
		}

	})

	t.Run("Ex 4 Printar cumprimentdo por periodo", func(t *testing.T) {
		date := time.Date(2019, 6, 5, 11, 51, 04, 0, time.UTC)
		cumpr := SayHelloPeriod("Viviane", date)
		expected := "Bom Dia Viviane"

		if cumpr != expected {
			t.Errorf("Expected %s, got %s", expected, cumpr)
		}
	})

	t.Run("Ex extra 1", func(t *testing.T) {
		list, sum := ModifyList(2, 3, 6)
		expected := 10

		if sum != expected {
			t.Errorf("Expected %v, got %v", expected, sum)
		}

		if list[0] != 1 {
			t.Errorf("Expected %v, got %v", 1, list[0])
		}

	})

	t.Run("Ex extra 2", func(t *testing.T) {
		num := CheckText("Estudar Go é muito bom", "g")

		if num != 1 {
			t.Errorf("Expected 1, got %v", num)
		}
	})

	t.Run("Ex extra 3", func(t *testing.T) {
		vendas := []Venda{
			{
				Nome:  "Tete",
				Valor: 653,
				Data:  time.Date(2019, 6, 5, 11, 51, 04, 0, time.UTC),
			},
			{
				Nome:  "TiTI",
				Valor: 658,
				Data:  time.Date(2019, 6, 8, 21, 51, 04, 0, time.UTC),
			},
			{
				Nome:  "Eita",
				Valor: 100,
				Data:  time.Date(2019, 6, 8, 21, 51, 04, 0, time.UTC),
			},
			{
				Nome:  "Eita",
				Valor: 200,
				Data:  time.Date(2019, 6, 8, 21, 51, 04, 0, time.UTC),
			},
			{
				Nome:  "Gigi",
				Valor: 741,
				Data:  time.Now(),
			},
			{
				Nome:  "Gigi",
				Valor: 100,
				Data:  time.Now(),
			},
		}

		result := ProcessSales(vendas)

		if result["Eita"] != 300 {
			t.Errorf("Expected %v, got %v", 300, result["Eita"])
		}
	})

	t.Run("Ex Extra 4", func(t *testing.T) {
		strForm := FormatWordForward("abc")

		if strForm != "def" {
			t.Errorf("Expected def, got %s", strForm)
		}

		strBack := FormatWordBack(strForm)

		if strBack != "abc" {
			t.Errorf("Expected abc, got %s", strBack)
		}
	})

}
