package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripleArgument(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hello world")

	tripleArgument(1, 2, "Hello")

	value := returnValue(2)
	fmt.Println("Value:", value)

	value1, value2 := doubleReturn(2)
	fmt.Println("value 1 y value 2:", value1, value2)

	value3, _ := doubleReturn(4)
	fmt.Println("value 3:", value3)

	anonFunction := func(a int) int {
		return a
	}
	a := anonFunction(10)
	fmt.Println("a:", a)
}
