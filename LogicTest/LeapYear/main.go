package main

import (
	"fmt"
	"strings"
)

func main() {
	var result []int

	for i := 1900; i <= 2020; i += 4 {
		if i%400 == 0 {
			result = append(result, i)
		} else if i%4 == 0 && i%100 != 0 {
			result = append(result, i)
		}
	}
	fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(result), " ", ", ", -1), "[]"))
}
