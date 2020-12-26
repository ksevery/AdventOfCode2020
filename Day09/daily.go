package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	lines := strings.Split(string(data), "\r\n")
	preamble := make([]int, 0, 25)

	for i := 0; i < 25; i++ {
		val, _ := strconv.Atoi(lines[i])
		preamble = append(preamble, val)
	}

	invalidVal := 0

	for i := 25; i < len(lines); i++ {
		val, _ := strconv.Atoi(lines[i])
		possibleCombinationsMap := make(map[int]int)
		for j := i - 25; j < len(preamble); j++ {
			firstVal := preamble[j]
			for ind := i - 25; ind < len(preamble); ind++ {
				if ind != j {
					secondVal := preamble[ind]
					possibleCombinationsMap[firstVal+secondVal] = 1
				}
			}
		}

		if _, ok := possibleCombinationsMap[val]; !ok {
			invalidVal = val
			break
		}

		preamble = append(preamble, val)
	}

	fmt.Println(invalidVal)
}
