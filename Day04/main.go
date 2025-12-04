package Day04

import "fmt"

func calcAdjacents(data []string, x int, y int) int {
	adjacents := 0

	// top
	if y > 0 && string(data[y-1][x]) == "@" {
		adjacents++
	}
	// bottom
	if y < len(data)-1 && string(data[y+1][x]) == "@" {
		adjacents++
	}
	// left
	if x > 0 && string(data[y][x-1]) == "@" {
		adjacents++
	}
	// right
	if x < len(data[y])-1 && string(data[y][x+1]) == "@" {
		adjacents++
	}
	// top left
	if y > 0 && x > 0 && string(data[y-1][x-1]) == "@" {
		adjacents++
	}
	// top right
	if y > 0 && x < len(data[y])-1 && string(data[y-1][x+1]) == "@" {
		adjacents++
	}
	// bottom left
	if y < len(data)-1 && x > 0 && string(data[y+1][x-1]) == "@" {
		adjacents++
	}
	// bottom right
	if y < len(data)-1 && x < len(data[y])-1 && string(data[y+1][x+1]) == "@" {
		adjacents++
	}

	return adjacents
}

func ResolvePart1(data []string) int {
	validRolls := 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if string(data[y][x]) != "@" {
				continue
			}

			if calcAdjacents(data, x, y) < 4 {
				validRolls++
			}
		}
	}
	return validRolls
}

func ResolvePart2(data []string) int {
	validRolls := 0
	for {
		validRollsThisSteps := 0
		for y := 0; y < len(data); y++ {
			for x := 0; x < len(data[y]); x++ {
				if string(data[y][x]) != "@" {
					continue
				}

				if calcAdjacents(data, x, y) < 4 {
					validRollsThisSteps++
					data[y] = fmt.Sprintf("%vx%v", data[y][:x], data[y][x+1:])
				}
			}
		}

		if validRollsThisSteps == 0 {
			break
		}

		validRolls += validRollsThisSteps
	}

	return validRolls
}

func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
