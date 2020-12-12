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
	for i := 0; i < len(splits); i++ {
		iVal, _ := strconv.ParseInt(splits[i], 10, 0)
		for j := 0; j < len(splits); j++ {
			if i == j {
				continue
			}

			jVal, _ := strconv.ParseInt(splits[j], 10, 0)
			if iVal+jVal <= 2020 {
				for n := 0; n < len(splits); n++ {
					if n == j || n == i {
						continue
					}

					nVal, _ := strconv.ParseInt(splits[n], 10, 0)
					if iVal+jVal+nVal == 2020 {
						fmt.Println(iVal * jVal * nVal)
						return
					}
				}
			}
		}
	}
}
