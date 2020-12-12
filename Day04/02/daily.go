package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("../input.txt")
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
			val, ok := passPartsMap[passPart]
			if !ok {
				isPassValid = false
				break
			} else if !validate(passPart, val) {
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

func validate(key string, val string) bool {
	switch key {
	case "byr":
		valNum, err := strconv.ParseInt(val, 10, 0)
		return len(val) == 4 && err == nil && valNum >= 1920 && valNum <= 2002
	case "iyr":
		valNum, err := strconv.ParseInt(val, 10, 0)
		return len(val) == 4 && err == nil && valNum >= 2010 && valNum <= 2020
	case "eyr":
		valNum, err := strconv.ParseInt(val, 10, 0)
		return len(val) == 4 && err == nil && valNum >= 2020 && valNum <= 2030
	case "hgt":
		if strings.HasSuffix(val, "cm") {
			cmVal := strings.Replace(val, "cm", "", 1)
			cmValNum, err := strconv.ParseInt(cmVal, 10, 0)
			return err == nil && cmValNum >= 150 && cmValNum <= 193
		} else if strings.HasSuffix(val, "in") {
			inVal := strings.Replace(val, "in", "", 1)
			inValNum, err := strconv.ParseInt(inVal, 10, 0)
			return err == nil && inValNum >= 59 && inValNum <= 76
		}

	case "hcl":
		if strings.HasPrefix(val, "#") {
			regex := regexp.MustCompile(`^#[0123456789abcdef]{6}$`)
			return regex.MatchString(val)
		}
	case "ecl":
		return val == "amb" || val == "blu" || val == "brn" || val == "gry" || val == "grn" || val == "hzl" || val == "oth"
	case "pid":
		_, err := strconv.ParseInt(val, 10, 0)
		return err == nil && len(val) == 9
	}

	return false
}
