package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	rows := strings.Split(string(data), "\r\n")

	maxID := 0.0
	ids := make(map[float64]struct{})
	for _, row := range rows {
		maxRow := 127.0
		minRow := 0.0

		minCol := 0.0
		maxCol := 7.0
		for i, char := range row {
			if i < 7 {
				if string(char) == "F" {
					maxRow = math.Floor(maxRow - ((maxRow - minRow) / 2.0))
				} else {
					minRow = math.Ceil(maxRow - ((maxRow - minRow) / 2))
				}
			} else {
				if string(char) == "L" {
					maxCol = math.Floor(maxCol - ((maxCol - minCol) / 2))
				} else {
					minCol = math.Ceil(maxCol - ((maxCol - minCol) / 2))
				}
			}
		}

		id := maxRow*8 + maxCol
		ids[id] = struct{}{}
		if id > maxID {
			maxID = id
		}
	}

	for i := 0.0; i < maxID; i++ {
		if _, ok := ids[i]; !ok {
			fmt.Println(i)
		}
	}

	fmt.Println(maxID)
}
