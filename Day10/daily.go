package main

import (
	"fmt"
	"io/ioutil"
	"sort"
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

	adapters := make([]int, 0, len(lines))

	for _, line := range lines {
		val, _ := strconv.Atoi(line)
		adapters = append(adapters, val)
	}

	sort.Ints(adapters)
	lastAdapter := 0
	oneJoltDifferences := 0
	threeJoltDifferences := 0

	for _, adapter := range adapters {
		difference := adapter - lastAdapter
		if difference == 1 {
			oneJoltDifferences++
		} else if difference == 3 {
			threeJoltDifferences++
		}

		lastAdapter = adapter
	}

	// Add built-in adapter difference
	threeJoltDifferences++
	fmt.Println(oneJoltDifferences * threeJoltDifferences)
}
