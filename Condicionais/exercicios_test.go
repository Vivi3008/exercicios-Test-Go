package condicionais

import "testing"

func TestExerciciosCondicionais(t *testing.T) {
	t.Run("Ex. 1: Pode votar?", func(t *testing.T) {
		anoNascimento := 1986
		maior := MaiorIdade(anoNascimento)

		if !maior {
			t.Errorf("Não pode votar")
		}
	})

	t.Run("Ex. 2 Numero positivo ou negativo", func(t *testing.T) {
		numero := Positivo(8)

		if !numero {
			t.Errorf("Número negativo")
		}
	})

	t.Run("Ex. 3 Verificar idade de uma pessoa", func(t *testing.T) {
		ageText := "Maior de idade"
		rangeAge := VerifyAge(36)

		if rangeAge != ageText {
			t.Errorf("Age expected %s, got %s", ageText, rangeAge)
		}
	})

}
