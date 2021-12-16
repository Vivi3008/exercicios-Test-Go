package lacosRepeticao

func NumberSequence(begin int, end int) []int {
	numbers := make([]int, 0)

	for i := begin; i <= end; i++ {
		numbers = append(numbers, i)
	}

	return numbers
}

func DayHours() []int {
	hours := make([]int, 0)

	for i := 0; i <= 24; {
		hours = append(hours, i)
		i++
	}

	return hours
}
