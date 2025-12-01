package Day01

import (
	"adventOfCode/utils"
	"strconv"
)

func ResolvePart1(data []string) int {

	var list1 []int
	var list2 []int

	for i := 0; i < len(data); i++ {
		splitted := utils.RegSplit(data[i], "\\s+")

		if value, err := strconv.Atoi(splitted[0]); err == nil {
			list1 = append(list1, value)
		}

		if value, err := strconv.Atoi(splitted[1]); err == nil {
			list2 = append(list2, value)
		}
	}

	return 0
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
