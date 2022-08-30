package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	textLower := strings.ToLower(text)
	var textReverse string

	for i := len(textLower) - 1; i >= 0; i-- {
		textReverse += string(textLower[i])
	}

	if textReverse == textLower {
		fmt.Println("Es palindromo")
	} else {
		fmt.Println("No es palindromo")
	}
}

func main() {
	// slice := []string{"hola", "que", "hace"}

	// for _, valor := range slice {
	// 	fmt.Println(valor)
	// }

	isPalindromo("Ama")
}
