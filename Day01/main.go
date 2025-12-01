package Day01

import (
	"adventOfCode/utils"
	"math"
	"strconv"
)

func ResolvePart1(data []string) int {
	countZeros := 0
	pos := 50

	for i := 0; i < len(data); i++ {
		lettre := data[i][:1]
		nombreStr := data[i][1:]

		nombre, err := strconv.Atoi(nombreStr)

		if err != nil {
			panic("aaaah")
		}

		if lettre == "R" {
			pos = (int)(math.Mod((float64)(pos+nombre), 100))
		} else if lettre == "L" {
			pos = (int)(math.Mod((float64)(pos-nombre), 100))
		}

		if pos == 0 {
			countZeros++
		}
	}

	return countZeros
}

func ResolvePart2(data []string) int {
	countZeros := 0
	pos := 50

	for i := 0; i < len(data); i++ {
		lettre := data[i][:1]
		nombreStr := data[i][1:]

		nombre, err := strconv.Atoi(nombreStr)

		if err != nil {
			panic("aaaah")
		}

		if len(nombreStr) > 2 {
			resteStr := nombreStr[len(nombreStr)-2:]
			reste, err := strconv.Atoi(resteStr)

			if err != nil {
				panic("aaaah")
			}

			centaine := (nombre - reste) / 100

			countZeros += centaine

			nombre = nombre - (centaine * 100)
		}

		previousPos := pos

		if lettre == "R" {
			pos = utils.Modulo(pos+nombre, 100)
			if (pos < previousPos) && previousPos != 0 {
				countZeros++
			}
		} else if lettre == "L" {
			pos = utils.Modulo(pos-nombre, 100)
			if (pos > previousPos || pos == 0) && previousPos != 0 {
				countZeros++
			}
		}
	}

	return countZeros
}

func Resolve(data []string) [2]any {
	return [2]any{
		ResolvePart1(data),
		ResolvePart2(data),
	}
}
