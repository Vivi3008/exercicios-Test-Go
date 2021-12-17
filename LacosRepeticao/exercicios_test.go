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

	t.Run("Ex 2,3 e 4. Imprimir horas do dia", func(t *testing.T) {
		hours, minutes := DayHours()

		if len(hours) != 24 {
			t.Errorf("Expected %v, got %v", 24, len(hours))
		}
		if len(minutes) != 1440 {
			t.Errorf("Expected %v, got %v", 1440, len(minutes))
		}

	})

	t.Run("Ex 5. Slice de strings", func(t *testing.T) {
		strings := SliceString()

		if len(strings) != 5 {
			t.Errorf("Expected %v, got %v", 5, len(strings))
		}
	})

	t.Run("Ex 6. Lista de compras", func(t *testing.T) {
		list := ShopList()

		if list[0] != "Arroz" {
			t.Errorf("Expected Arroz, got %v", list[0])
		}
	})
}
