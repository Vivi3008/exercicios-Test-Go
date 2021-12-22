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

}
