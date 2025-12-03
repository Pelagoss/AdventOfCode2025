package Day03

import (
	"bytes"
	"fmt"
	"strconv"
)

func ResolveParts(data []string, digitLength int) int {
	count := 0

	for _, bank := range data {
		var digits []int

		for i := 0; i < digitLength; i++ {
			digits = append(digits, 0)
		}

		for i := 0; i < len(bank); i++ {
			batterie, _ := strconv.Atoi(string(bank[i]))

			for j := 0; j < len(digits); j++ {
				if batterie > digits[j] && i < len(bank)-len(digits)+j+1 {
					for k := j + 1; k < len(digits); k++ {
						digits[k] = 0
					}
					digits[j] = batterie
					break
				}
			}
		}

		var buffer bytes.Buffer

		for _, digit := range digits {
			buffer.WriteString(fmt.Sprintf("%v", digit))
		}

		toAdd, _ := strconv.Atoi(buffer.String())

		count += toAdd
	}

	return count
}

func ResolvePart2(data []string) int {
	count := 0

	for _, bank := range data {
		digits := [12]int{}

		for i := 0; i < len(bank); i++ {
			batterie, _ := strconv.Atoi(string(bank[i]))

			for j := 0; j < len(digits); j++ {
				if batterie > digits[j] && i < len(bank)+j-1 {
					for k := j + 1; k < len(digits); k++ {
						digits[k] = 0
					}
					digits[j] = batterie
					break
				}
			}
		}

		var buffer bytes.Buffer

		for _, digit := range digits {
			buffer.WriteString(fmt.Sprintf("%v", digit))
		}

		toAdd, _ := strconv.Atoi(buffer.String())

		count += toAdd
	}

	return count
}

func Resolve(data []string) [2]any {
	return [2]any{
		ResolveParts(data, 2),
		ResolveParts(data, 12),
	}
}
