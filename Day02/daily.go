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

	splits := strings.Split(string(data), "\r\n")
	validPasswordsCount := 0

	for i := 0; i < len(splits); i++ {
		lineSplit := strings.Split(splits[i], " ")
		limits := lineSplit[0]
		letter := lineSplit[1][:len(lineSplit[1])-1]
		password := lineSplit[2]

		limitParts := strings.Split(limits, "-")
		minLetters, _ := strconv.ParseInt(limitParts[0], 10, 0)
		maxLetters, _ := strconv.ParseInt(limitParts[1], 10, 0)

		if minLetters == 0 {
			validPasswordsCount++
			continue
		} else if !strings.Contains(password, letter) {
			continue
		} else {
			letterInPasswordCount := strings.Count(password, letter)
			if letterInPasswordCount <= int(maxLetters) && letterInPasswordCount >= int(minLetters) {
				validPasswordsCount++
			}
		}
	}

	fmt.Println(validPasswordsCount)
}
