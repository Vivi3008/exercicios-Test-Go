package lacosRepeticao

import "testing"

func TestExerciciosLacosRepeticao(t *testing.T) {
	t.Run("Ex. 1 Imprimir numeros usando for", func(t *testing.T) {
		begin := 13
		end := 27

		result := NumberSequence(begin, end)

		if len(result) != 15 {
			t.Errorf("Expected %v, got %v", 15, len(result))
		}

		other := NumberSequence(1, 10)

		if len(other) != 10 {
			t.Errorf("Expected %v, got %v", 10, len(other))
		}
	})

	t.Run("Ex 2. Imprimir horas do dia", func(t *testing.T) {
		result := DayHours()

		if len(result) != 25 {
			t.Errorf("Expected %v, got %v", 25, len(result))
		}
	})
}
