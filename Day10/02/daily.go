package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

	adapters := make([]int, 0, len(lines))

	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		adapters = append(adapters, val)
	}

	sort.Ints(adapters)
	// currentAdapter := 0
	possibleWaysToArrange := 1 // there is always the default arrangement
	// waysToArrangeMap := make(map[int][]int)

	for i := 0; i < len(adapters); i++ {
		// difference := adapters[i] - currentAdapter
		innerArrangements := arrangementsStartingWithIndex(adapters, i, 0)

		possibleWaysToArrange += innerArrangements
		// currentAdapter = adapters[i]
	}

	// for _, arrangements := range waysToArrangeMap {
	// 	if len(arrangements) > 1 {
	// 		possibleWaysToArrange += len(arrangements)
	// 	}
	// }

	// Add built-in adapter difference
	fmt.Println(possibleWaysToArrange)
}

func arrangementsStartingWithIndex(adapters []int, index int, arrangementsCount int) int {
	currentAdapter := adapters[index]
	for i := index + 1; i < len(adapters); i++ {
		difference := adapters[i] - currentAdapter
		if difference <= 3 {
			arrangementsCount++
			innerCount := arrangementsStartingWithIndex(adapters, i, arrangementsCount)
			arrangementsCount += innerCount
		}
	}

	return arrangementsCount
}
