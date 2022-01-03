package metodos

import (
	"errors"
	"reflect"
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

		if sum != 21 {
			t.Errorf("Expected sum = 21, got %v", sum)
		}

		if media != 3.5 {
			t.Errorf("Expected media %v, got %v", 3.5, media)
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

		if remove[len(remove)-1] != 9 {
			t.Errorf("Expected last item 9, got %v", remove[len(remove)-1])
		}
	})

	t.Run("Ex extra 2", func(t *testing.T) {
		var cpf CPF = "11144477735"
		expected := []int{1, 1, 1, 4, 4, 4, 7, 7, 7, 3, 5}
		result, err := cpf.ValidaCpf()

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		if !reflect.DeepEqual(expected, result) {
			t.Errorf("Expected %v, got %v", expected, result)
		}
	})

	t.Run("Ex extra 2 Fail if cpf has wrong caracters", func(t *testing.T) {
		var cpf CPF = "a3255642153"
		result, err := cpf.ValidaCpf()

		if !errors.Is(err, ErrInvalidCaracters) {
			t.Errorf("Expected %s, got %s", ErrInvalidCaracters, err)
		}

		if len(result) != 0 {
			t.Errorf("Expected result be empty slice")
		}
	})

	t.Run("Ex extra 2 Fail if cpf is invalid", func(t *testing.T) {
		var cpf CPF = "11566588852"
		result, err := cpf.ValidaCpf()

		if !errors.Is(err, ErrInvalidCpf) {
			t.Errorf("Expected %s, got %s", ErrInvalidCpf, err)
		}

		if len(result) != 0 {
			t.Errorf("Expected result be empty slice")
		}
	})
}
