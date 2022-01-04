package interfaces

import (
	"reflect"
	"testing"
)

func TestExerciciosInterface(t *testing.T) {
	t.Run("Ex 1 interfaces", func(t *testing.T) {
		circ := Circulo{
			Raio: 4,
		}
		quad := Quadrado{
			Lado: 4,
		}
		areaCirc := CalculaAreaGeometrica(circ)

		areaQuad := CalculaAreaGeometrica(quad)

		var expected float64 = 50.265482

		if reflect.DeepEqual(areaCirc, expected) {
			t.Errorf("Expected %f, got %f", expected, areaCirc)
		}

		if areaQuad != 16 {
			t.Errorf("Expected %v, got %v", 16, areaQuad)
		}
	})
}
