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

	t.Run("Ex extra 1 Pilha", func(t *testing.T) {
		pilha := Pilha{1, 2}
		result := pilha.Push(6)

		if result[len(result)-1] != 6 {
			t.Errorf("Expected last item 6, got %v", result[len(result)-1])
		}

		p2 := Pilha{6, 5, 8, 9, 23}
		remove := p2.Pop()
		fmt.Println(remove)
		if remove[len(remove)-1] != 9 {
			t.Errorf("Expected last item 9, got %v", remove[len(remove)-1])
		}
	})
}
