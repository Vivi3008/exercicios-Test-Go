package funcoes

import (
	"testing"
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

}
