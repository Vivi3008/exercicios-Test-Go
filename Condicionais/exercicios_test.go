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

	t.Run("Ex. 4 Verificar idade usando Switch", func(t *testing.T) {
		ageText := "Idoso"
		rangeAge := VerifyAgeSwitch(85)

		if rangeAge != ageText {
			t.Errorf("Expected %s, got %s", ageText, rangeAge)
		}
	})

	t.Run("Ex. 5 Verificar período do dia", func(t *testing.T) {
		hour := 17
		period := VerifyDayPeriod(hour)

		if period != "Tarde" {
			t.Errorf("Expected %s, got %s", "Tarde", period)
		}
	})

	t.Run("Ex Extra 1 Maior valor", func(t *testing.T) {
		x := 9
		y := 8
		z := -10

		maior := CheckBigger(x, y, z)

		if maior != x {
			t.Errorf("Expected %v, got %v", x, maior)
		}
	})

	t.Run("Ex Extra 2 Multiplos", func(t *testing.T) {
		muliple := CheckMultiple(12)
		txt := "Multiplo de 2"

		if muliple != txt {
			t.Errorf("Expected %v, got %v", txt, muliple)
		}
	})
}
