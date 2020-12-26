package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("../input.txt")
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
	invalidValIndex := 0

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
			invalidValIndex = i
			break
		}

		preamble = append(preamble, val)
	}

	weaknessSum := 0

	currentSet := make([]int, 0, 1)
	haveFoundSet := false
	for j := 0; j < invalidValIndex; j++ {
		firstVal, _ := strconv.Atoi(lines[j])
		currSum := firstVal
		currentSet = append(currentSet, firstVal)
		for ind := j + 1; ind < invalidValIndex; ind++ {
			if ind != j {
				secondVal, _ := strconv.Atoi(lines[ind])
				currSum += secondVal
				if currSum == invalidVal {
					haveFoundSet = true
					currentSet = append(currentSet, secondVal)
					break
				} else if currSum > invalidVal {
					break
				} else {
					currentSet = append(currentSet, secondVal)
				}
			}
		}

		if haveFoundSet {
			break
		} else if currSum > invalidVal {
			currentSet = make([]int, 0, 1)
		}
	}

	if haveFoundSet {
		min, max := minMax(currentSet)
		weaknessSum = min + max
	}

	fmt.Println(weaknessSum)
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
