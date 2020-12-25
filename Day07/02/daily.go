package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var bagsContainingGoldStr []string = make([]string, 1, 1)
var firstLevelBags []string = make([]string, 1, 1)

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	lines := strings.Split(string(data), "\r\n")

	shinyGoldBagStr := "shiny gold bag"
	regex := regexp.MustCompile(`(\d )`)
	bagsMap := make(map[string]map[string]int)

	for _, line := range lines {
		lineParts := strings.Split(line, "contain")
		containingBag := lineParts[0]
		trimmedContainerBag := strings.TrimSuffix(strings.TrimSpace(containingBag), "s")
		contains := lineParts[1]
		bagsWithin := strings.Split(contains, ", ")
		if strings.HasPrefix(contains, "no") {
			continue
		}

		bagsMap[trimmedContainerBag] = make(map[string]int)
		for _, bag := range bagsWithin {
			trimmedBag := strings.TrimSpace(strings.TrimSuffix(bag, "."))
			trimmedBag = strings.TrimSuffix(trimmedBag, "s")
			bagsCount, _ := strconv.Atoi(strings.TrimSpace(regex.FindString(trimmedBag)))
			trimmedBag = regex.ReplaceAllString(trimmedBag, "")
			bagsMap[trimmedContainerBag][trimmedBag] = bagsCount
		}
	}

	shinyGoldBagContents := bagsMap[shinyGoldBagStr]

	bagsInShiny := searchRecursively(bagsMap, shinyGoldBagContents, 0)

	fmt.Println(bagsInShiny)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func searchRecursively(bagsMap map[string]map[string]int, bagsContain map[string]int, count int) int {
	for bag, bagCount := range bagsContain {
		multiplier := bagCount
		if bagMap, ok := bagsMap[bag]; ok {
			internalCount := searchRecursively(bagsMap, bagMap, 1)
			count += multiplier * internalCount
		} else if bag == "no other bag" {
			count = 1
		}
	}

	return count
}
