package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	groups := strings.Split(string(data), "\r\n\r\n")

	totalCount := 0

	for _, group := range groups {
		answersInGroup := strings.Split(group, "\r\n")
		answersMapForGroup := make(map[string]int)
		for _, personAnswers := range answersInGroup {
			for _, answer := range personAnswers {
				if _, ok := answersMapForGroup[string(answer)]; ok {
					answersMapForGroup[string(answer)]++
				} else {
					answersMapForGroup[string(answer)] = 1
				}
			}
		}

		fmt.Println(answersMapForGroup)
		for _, val := range answersMapForGroup {
			// fmt.Println(val)
			// fmt.Println(len(answersInGroup))
			if val >= len(answersInGroup) {
				totalCount += 1
			}
		}

		// fmt.Println()
	}

	fmt.Println(totalCount)
}
