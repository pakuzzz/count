package main

import (
	"fmt"
)

func isVowel(r rune) bool {
	vowels := []rune{'а', 'е', 'ё', 'и', 'о', 'у', 'ы', 'э', 'ю', 'я'}
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}

func main() {
	var text string
	fmt.Print("Введите строку (без пробелов и только строчные буквы): ")
	fmt.Scanln(&text)

	count := 0
	for _, r := range text {
		if isVowel(r) {
			count++
		}
	}

	fmt.Println("Количество гласных:", count)
}
