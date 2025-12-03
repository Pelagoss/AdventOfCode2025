package Day03

import (
	"fmt"
	"strconv"
)

func ResolvePart1(data []string) int {
	count := 0

	for _, bank := range data {
		higher := 0
		secondHigher := 0

		for i := 0; i < len(bank); i++ {
			batterie, _ := strconv.Atoi(string(bank[i]))

			if batterie > higher && i != len(bank)-1 {
				secondHigher = 0
				higher = batterie
			} else if batterie > secondHigher {
				secondHigher = batterie
			}
		}

		toAdd, _ := strconv.Atoi(fmt.Sprintf("%v%v", higher, secondHigher))

		count += toAdd
	}

	return count
}

func ResolvePart2(data []string) int {
	return 0
}

func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
