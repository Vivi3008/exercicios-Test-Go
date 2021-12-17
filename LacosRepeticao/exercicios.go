package lacosRepeticao

func NumberSequence(begin int, end int) []int {
	numbers := make([]int, 0)

	for i := begin; i <= end; i++ {
		numbers = append(numbers, i)
	}

	return numbers
}

func DayHours() ([]int, []int) {
	hours := make([]int, 0)
	minutes := make([]int, 0)

	for i := 0; i < 24; {
		hours = append(hours, i)
		for m := 0; m < 60; m++ {
			minutes = append(minutes, m)
		}
		i++
	}

	return hours, minutes
}

func SliceString() []string {
	strings := []string{"um", "dois", "tres", "quatro", "cinco"}

	return strings
}

func ShopList() []string {
	list := []string{"Arroz", "Carne", "Chocolate", "Chicletes", "Verduras", "Refrigerante"}

	return list
}
