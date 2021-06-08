package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	s := strings.Fields(input)

	var result []string

	for i := 0; i < len(s); i++ {
		rns := []rune(s[i])
		if unicode.IsUpper(rns[0]) {
			rns[len(rns)-1] = unicode.ToUpper(rns[len(rns)-1])
		}
		for x, j := 0, len(rns)-1; x < j; x, j = x+1, j-1 {
			rns[x], rns[j] = rns[j], rns[x]
		}
		result = append(result, string(rns))
	}
	fmt.Println(strings.Join(result, " "))
}
