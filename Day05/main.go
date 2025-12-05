package Day05

import (
	"adventOfCode/utils"
	"fmt"
	"strconv"
)

func ResolvePart1(data []string) int {
	fresh := 0
	var freshRanges [][2]int
	isIngredient := false

	for i := 0; i < len(data); i++ {
		if isIngredient == false {
			if data[i] == "" {
				isIngredient = true
				continue
			}

			split := utils.RegSplit(data[i], "-")

			startRange, _ := strconv.Atoi(split[0])
			endRange, _ := strconv.Atoi(split[1])

			freshRanges = append(freshRanges, [2]int{startRange, endRange})
		} else {
			id, _ := strconv.Atoi(data[i])

			for j := 0; j < len(freshRanges); j++ {
				if id >= freshRanges[j][0] && id <= freshRanges[j][1] {
					fresh++
					break
				}
			}
		}
	}

	return fresh
}

func ResolvePart2(data []string) int {
	isIngredient := false
	var freshRanges [][2]int

	for i := 0; i < len(data); i++ {
		if isIngredient == true {
			break
		}

		if data[i] == "" {
			isIngredient = true
			continue
		}

		split := utils.RegSplit(data[i], "-")

		for {
			startRange, _ := strconv.Atoi(split[0])
			endRange, _ := strconv.Atoi(split[1])

			var toDelete []int
			continued := true

			for j := 0; j < len(freshRanges); j++ {
				startInRange := startRange >= freshRanges[j][0] && startRange <= freshRanges[j][1]
				endInRange := endRange >= freshRanges[j][0] && endRange <= freshRanges[j][1]
				if startInRange == false && endInRange == false {
					continue
				}
				continued = false

				if startInRange {
					startRange = min(freshRanges[j][0], startRange)
				}

				if endInRange {
					endRange = max(freshRanges[j][1], endRange)
				}

				toDelete = append(toDelete, j)
			}

			for l := 0; l < len(toDelete); l++ {
				freshRanges[toDelete[l]][0] = 0
				freshRanges[toDelete[l]][1] = 0
			}

			freshRanges = append(freshRanges, [2]int{startRange, endRange})

			if continued == false {
				break
			}
		}
	}

	freshIds := 0
	for _, freshRange := range freshRanges {
		if freshRange[0] == 0 && freshRange[1] == 0 {
			continue
		}

		fmt.Println(freshRange)

		freshIds += (freshRange[1] - freshRange[0]) + 1
	}

	return freshIds
}

func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
