package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	requiredPassParts := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	passports := strings.Split(string(data), "\r\n\r\n")
	validPasses := 0

	for _, passport := range passports {
		passParts := strings.FieldsFunc(passport, split)
		passPartsMap := make(map[string]string)
		for _, part := range passParts {
			singlePassPart := strings.Split(part, ":")
			passPartsMap[singlePassPart[0]] = singlePassPart[1]
		}

		isPassValid := true
		for _, passPart := range requiredPassParts {
			if _, ok := passPartsMap[passPart]; !ok {
				isPassValid = false
				break
			}
		}

		if isPassValid {
			validPasses++
		}
	}

	fmt.Println(validPasses)
}

func split(r rune) bool {
	isSpace := r == ' '
	isCarriageReturn := r == 13 || r == 10 // Carriage return and linefeed, \r\n
	return isSpace || isCarriageReturn
}
