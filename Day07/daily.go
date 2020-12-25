package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var bagsContainingGoldStr []string = make([]string, 1, 1)
var firstLevelBags []string = make([]string, 1, 1)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	lines := strings.Split(string(data), "\r\n")

	bagsContainingGold := 0
	shinyGoldBag := "shiny gold bag"
	bagsMap := make(map[string][]string)
	regex := regexp.MustCompile(`(\d )`)

	for _, line := range lines {
		lineParts := strings.Split(line, "contain")
		containingBag := lineParts[0]
		trimmedContainerBag := strings.TrimSuffix(strings.TrimSpace(containingBag), "s")
		contains := lineParts[1]
		bagsWithin := strings.Split(contains, ", ")
		if strings.HasPrefix(contains, "no") {
			continue
		}

		bagsMap[trimmedContainerBag] = make([]string, 1, 1)
		for _, bag := range bagsWithin {
			trimmedBag := strings.TrimSpace(strings.TrimSuffix(bag, "."))
			trimmedBag = strings.TrimSuffix(regex.ReplaceAllString(trimmedBag, ""), "s")
			bagsMap[trimmedContainerBag] = append(bagsMap[trimmedContainerBag], trimmedBag)
		}
	}

	for key, val := range bagsMap {
		if contains(val, shinyGoldBag) {
			bagsContainingGold++
			bagsContainingGoldStr = append(bagsContainingGoldStr, key)
			firstLevelBags = append(firstLevelBags, key)
		} else {

			cointainsRecursively := searchRecursively(bagsMap, shinyGoldBag, val)
			if cointainsRecursively {
				bagsContainingGold++
				bagsContainingGoldStr = append(bagsContainingGoldStr, key)
			}
		}
	}

	fmt.Println(bagsContainingGold)
	fmt.Println(bagsContainingGoldStr)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func searchRecursively(bagsMap map[string][]string, textToFind string, bagsContain []string) bool {
	for _, bag := range bagsContain {
		if bagMap, ok := bagsMap[bag]; ok {
			if contains(bagMap, textToFind) {
				return true
			} else {
				if searchRecursively(bagsMap, textToFind, bagMap) {
					return true
				}
			}
		}
	}

	return false
}
