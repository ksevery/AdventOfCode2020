package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const treeStr = "#"

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	treeCount11 := 0
	treeCount31 := 0
	treeCount51 := 0
	treeCount71 := 0
	treeCount12 := 0

	rows := strings.Split(string(data), "\r\n")

	treeCount11 = getTreeCount(rows, 1, 1)
	treeCount12 = getTreeCount(rows, 2, 1)
	treeCount31 = getTreeCount(rows, 1, 3)
	treeCount51 = getTreeCount(rows, 1, 5)
	treeCount71 = getTreeCount(rows, 1, 7)

	fmt.Printf("Trees in path 1, 1: %v \n", treeCount11)
	fmt.Printf("Trees in path 1, 2: %v \n", treeCount12)
	fmt.Printf("Trees in path 3, 1: %v \n", treeCount31)
	fmt.Printf("Trees in path 5, 1: %v \n", treeCount51)
	fmt.Printf("Trees in path 7, 1: %v \n", treeCount71)

	multipliedTrees := treeCount11 * treeCount12 * treeCount31 * treeCount51 * treeCount71
	fmt.Println(multipliedTrees)
}

func getTreeCount(rows []string, rowsDown int, colsRight int) int {
	rowLength := len(rows[0])
	downPathLength := len(rows) - 1
	rightPathLength := downPathLength*colsRight + 1
	copyCount := rightPathLength/rowLength + rowsDown
	treesCount := 0

	for i := rowsDown; i < len(rows); i += rowsDown {
		fullRow := strings.Repeat(rows[i], copyCount)
		charAtPos := string(fullRow[i*colsRight])
		if charAtPos == treeStr {
			treesCount++
		}
	}

	return treesCount
}
