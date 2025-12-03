package Day02

import (
	"adventOfCode/utils"
	"fmt"
	"math"
	"strconv"
)

func ResolvePart1(data []string) int {
	count := 0
	idRanges := utils.RegSplit(data[0], ",")

	for _, idRange := range idRanges {
		bornes := utils.RegSplit(idRange, "-")

		startId := bornes[0]
		endId := bornes[1]

		if len(startId)%2 != 0 && len(endId)%2 != 0 {
			continue
		}

		currentId, err := strconv.Atoi(startId)
		finishId, err := strconv.Atoi(endId)

		if err != nil {
			panic("aaaah")
		}

		if len(startId)%2 != 0 {
			currentId = int(math.Pow10(len(startId)))
		}

		if len(endId)%2 != 0 {
			finishId = int(math.Pow10(len(endId)-1)) - 1
		}

		for {
			currentIdStr := fmt.Sprintf("%v", currentId)

			if currentIdStr[:len(currentIdStr)/2] == currentIdStr[len(currentIdStr)/2:] {
				count += currentId
			}

			if currentId == finishId {
				break
			}

			currentId++
		}
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
