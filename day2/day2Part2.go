package day2

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func PossibleGamePart2() {
	file, _ := os.Open("day2/input.txt")

	inputByte, _ := io.ReadAll(file)

	input := string(inputByte)

	splited := strings.Split(input, "\n")

	m := map[string]int{"red": 0, "green": 0, "blue": 0}

	val := 0

	for _, game := range splited {
		gameSplited := strings.FieldsFunc(game, Split)
		// valid := true
		for _, v := range gameSplited {
			tempSplited := strings.FieldsFunc(v, Split2)
			for i := 0; i <= len(tempSplited)-2; i += 2 {
				currentVal, currentType := tempSplited[i], tempSplited[i+1]
				currentValInt, _ := strconv.Atoi(currentVal)

				if currentValInt > m[currentType] {
					m[currentType] = currentValInt
				}
			}

		}
		val += m["red"] * m["green"] * m["blue"]
		m = map[string]int{"red": 0, "green": 0, "blue": 0}
	}

	fmt.Println(val)
}
