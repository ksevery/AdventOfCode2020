package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var lastJumpInstruction int = 0
var accumulator int = 0

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	lines := strings.Split(string(data), "\r\n")
	didProgramTerminate := false

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		lineParts := strings.Split(line, " ")
		mainCommand := lineParts[0]
		val, _ := strconv.Atoi(lineParts[1])
		if mainCommand == "jmp" {
			mainCommand = "nop"
		} else if mainCommand == "nop" && val != 0 {
			mainCommand = "jmp"
		}

		localAccumulator := 0
		executionMap := make(map[int]int)

		for j := 0; j < len(lines); j++ {
			line := lines[j]
			lineParts := strings.Split(line, " ")
			command := lineParts[0]
			val, _ := strconv.Atoi(lineParts[1])
			if j == i {
				command = mainCommand
			}

			if _, ok := executionMap[j]; ok {
				didProgramTerminate = true
				break
			} else {
				executionMap[j] = 1

				switch command {
				case "nop":
					break
				case "acc":
					localAccumulator += val
					break
				case "jmp":
					j += val - 1
					lastJumpInstruction = val
					break
				}
			}
		}

		if !didProgramTerminate {
			accumulator = localAccumulator
			break
		} else {
			didProgramTerminate = false
		}
	}

	if didProgramTerminate {
		fmt.Println("Program terminated, find another solution")
	}

	fmt.Println(accumulator)
}
