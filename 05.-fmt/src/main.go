package main

import "fmt"

func main() {
	helloMessage := "Hello"
	worlMessage := "world"

	// Println
	fmt.Println(helloMessage, worlMessage)
	fmt.Println(helloMessage, worlMessage)

	// Printf
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)
	// Usar %v cuando no conocemos los tipos de datos
	fmt.Printf("%v is %v years old.\n", name, age)

	// Sprintf
	message := fmt.Sprintf("%v is %v years old.", name, age)
	fmt.Println(message)

	// Tipo de dato
	fmt.Printf("helloMessage: %T\n", helloMessage)
	fmt.Printf("age: %T", age)
}
