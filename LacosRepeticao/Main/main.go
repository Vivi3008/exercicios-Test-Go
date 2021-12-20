package main

import (
	"fmt"

	lacosRepeticao "github.com/Vivi3008/exercicios-Test-Go/LacosRepeticao"
)

func main() {
	//ex 1
	fmt.Printf("Sequencia de numeros %v\n", lacosRepeticao.NumberSequence(1, 30))
	// ex 2, 3 e 4
	hours, minutes := lacosRepeticao.DayHours()

	for _, hour := range hours {
		for _, minute := range minutes {
			fmt.Printf("%.2d:%.2d\n", hour, minute)
		}
	}

	//ex 5
	strings := lacosRepeticao.SliceString()

	for index, str := range strings {
		fmt.Println(index, str)
	}

	//ex 6
	list := lacosRepeticao.ShopList()

	for i, v := range list {
		fmt.Println(i+1, v)
	}

	//ex extra1
	for i := 0; i < len(list); i++ {
		fmt.Println(i+1, list[i])
	}
	//ex extra 3
	lacosRepeticao.PrintValues()
}
