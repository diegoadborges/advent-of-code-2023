package day1

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func SumFirstNumAndLast() {
	file, err := os.Open("day1/input.txt")
	check(err)

	inputByte, err := io.ReadAll(file)

	input := string(inputByte)

	sum := 0
	splited := strings.Split(input, "\n")
	for _, v := range splited {
		var first, last string
		current := strings.Trim(v, "	")
		for i := 0; i < len(current); i++ {

			if current[i] >= 38 && current[i] <= 57 {
				if first == "" {
					first = string(current[i])
					continue
				}
				last = string(current[i])
			}
		}
		if last == "" {
			last = first
		}
		val, _ := strconv.Atoi(first + last)
		sum += val
	}
	fmt.Println(sum)
}
