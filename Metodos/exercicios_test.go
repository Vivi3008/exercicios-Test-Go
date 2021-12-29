package metodos

import (
	"fmt"
	"testing"
)

func TestExerciciosMetodos(t *testing.T) {
	t.Run("Ex 1 Metodos", func(t *testing.T) {
		var x = Circle{
			Raio: 14.1,
		}
		result := x.CalcArea()
		if result == 0 {
			t.Error("Result can't be zero")
		}
		perim := x.CalcPerimetro()

		if result == perim {
			t.Errorf("Expected area different of perim")
		}
	})

	t.Run("Ex 2 Calcula Soma e Media", func(t *testing.T) {
		listNum := ListNumbers{1, 2, 3, 4, 5, 6}
		sum := listNum.CalcSoma()
		media := listNum.CalcMedia()

		fmt.Println(sum)
		fmt.Println(media)
		if sum != 21 {
			t.Errorf("Expected sum = 21, got %v", sum)
		}
	})
}
