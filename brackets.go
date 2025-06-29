package main

import (
	"fmt"
)

func main() {
	var formula string
	fmt.Print("Введите строку (без пробелов): ")
	fmt.Scanln(&formula)

	open := 0
	close := 0

	for _, r := range formula {
		if r == '(' {
			open++
		}
		if r == ')' {
			close++
		}
	}

	if open == close {
		fmt.Printf("Скобки расставлены верно, %d открывающиеся, %d закрывающиеся\n", open, close)
	} else {
		fmt.Printf("Скобки расставлены неправильно, %d открывающиеся, %d закрывающиеся\n", open, close)
	}
}
