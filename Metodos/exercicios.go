package metodos

import "math"

// ex 1
type Circle struct {
	Raio float64
}

func (c Circle) CalcArea() float64 {
	area := math.Pi * math.Pow(c.Raio, 2)

	return area
}

func (c Circle) CalcPerimetro() float64 {
	perim := 2 * (math.Pi * c.Raio)

	return perim
}

//ex 2
type ListNumbers []int

func (l ListNumbers) CalcSoma() int {
	var count int
	for _, v := range l {
		count += v
	}
	return count
}

func (l ListNumbers) CalcMedia() float64 {
	media := float64(l.CalcSoma()) / float64(len(l))

	return media
}
