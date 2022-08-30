package main

import "fmt"

func main() {
	switch module := 5 % 2; module {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condicion
	value := 200
	switch {
	case value > 100:
		fmt.Println("Es mayor a 100")
	case value < 0:
		fmt.Println("Es mayor a 100")
	default:
		fmt.Println("No condicion")
	}

}
