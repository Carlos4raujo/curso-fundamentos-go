package main

import "fmt"

func main() {
	// constants

	const pi float64 = 3.14
	const pi2 = 3.1415
	fmt.Println("pi:", pi)
	fmt.Println("pi2:", pi2)

	//variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	// Zero values
	var a int     // integers 		-> 0
	var b float64 // decimals		-> 0
	var c string  // text 			-> ""
	var d bool    // true or false	-> false

	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	const areaCuadrado = baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)
}
