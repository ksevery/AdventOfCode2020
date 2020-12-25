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

	lines := strings.Split(string(data), "\r\n")
	executionMap := make(map[int]int)
	accumulator := 0

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if _, ok := executionMap[i]; ok {
			break
		} else {
			executionMap[i] = 1
			lineParts := strings.Split(line, " ")
			switch lineParts[0] {
			case "nop":
				break
			case "acc":
				val, _ := strconv.Atoi(lineParts[1])
				accumulator += val
				break
			case "jmp":
				val, _ := strconv.Atoi(lineParts[1])
				i += val - 1
				break
			}
		}
	}

	fmt.Println(accumulator)
}
