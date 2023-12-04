package day2

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func PossibleGame() {
	file, _ := os.Open("day2/input.txt")

	inputByte, _ := io.ReadAll(file)

	input := string(inputByte)

	splited := strings.Split(input, "\n")

	m := map[string]int{"red": 12, "green": 13, "blue": 14}

	val := 0

	for _, game := range splited {
		gameSplited := strings.FieldsFunc(game, Split)
		valid := true
		for _, v := range gameSplited {
			tempSplited := strings.FieldsFunc(v, Split2)

			for i := 0; i <= len(tempSplited)-2; i += 2 {
				currentVal, currentType := tempSplited[i], tempSplited[i+1]
				currentValInt, _ := strconv.Atoi(currentVal)
				fmt.Println(currentValInt, currentType)
				m[currentType] -= currentValInt
			}

			fmt.Println("-")

			if m["red"] < 0 || m["green"] < 0 || m["blue"] < 0 {
				// fmt.Println(gameSplited[0])
				valid = false
			}

			m["red"] = 12
			m["green"] = 13
			m["blue"] = 14
		}
		if valid {
			strId := strings.Split(gameSplited[0], " ")[1]
			id, _ := strconv.Atoi(strId)
			// fmt.Println(id)
			val += id
		}
	}

	fmt.Println(val)
}

func Split(r rune) bool {
	return r == ':' || r == ';'
}

func Split2(r rune) bool {
	return r == ',' || r == ' '
}
