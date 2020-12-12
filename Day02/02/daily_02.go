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

	splits := strings.Split(string(data), "\r\n")
	validPasswordsCount := 0

	for i := 0; i < len(splits); i++ {
		lineSplit := strings.Split(splits[i], " ")
		limits := lineSplit[0]
		letter := lineSplit[1][:len(lineSplit[1])-1]
		password := lineSplit[2]

		limitParts := strings.Split(limits, "-")
		firstIndex64, _ := strconv.ParseInt(limitParts[0], 10, 0)
		secondIndex64, _ := strconv.ParseInt(limitParts[1], 10, 0)

		firstIndex := int(firstIndex64) - 1
		lastIndex := int(secondIndex64) - 1

		if !strings.Contains(password, letter) {
			continue
		} else if len(password) > lastIndex {
			charAtFirstIndex := string(password[firstIndex])
			charAtLastIndex := string(password[lastIndex])
			if charAtFirstIndex == letter && charAtLastIndex != letter {
				validPasswordsCount++
			} else if charAtFirstIndex != letter && charAtLastIndex == letter {
				validPasswordsCount++
			}
		} else if string(password[firstIndex]) == letter {
			validPasswordsCount++
		}
	}

	fmt.Println(validPasswordsCount)
}
