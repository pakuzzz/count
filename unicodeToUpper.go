package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func capitalize(word string) string {
	runes := []rune(word)
	for i := range runes {
		if i == 0 {
			runes[i] = unicode.ToUpper(runes[i])
		} else {
			runes[i] = unicode.ToLower(runes[i])
		}
	}
	return string(runes)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	words := strings.Fields(input)
	for i, word := range words {
		words[i] = capitalize(word)
	}

	result := strings.Join(words, " ")
	fmt.Println(result)
}
