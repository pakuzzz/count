package main

import (
	"fmt"
)

func main() {
	var text string
	fmt.Print("Введите строку: ")
	fmt.Scanln(&text)

	runes := []rune(text)
	fmt.Println("Количество символов:", len(runes))
}
