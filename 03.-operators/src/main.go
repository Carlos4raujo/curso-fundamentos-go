package main

import "fmt"

func main() {
	const baseCuadrado = 10
	const areaCuadrado = baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// Suma
	result := x + y
	fmt.Println("Suma: ", result)

	// Resta
	result = y - x
	fmt.Println("Resta: ", result)

	// Multiplicacion
	result = x * y
	fmt.Println("Multiplicación: ", result)

	// División
	result = y / x
	fmt.Println("División: ", result)

	// Residuo
	result = y % x
	fmt.Println("Residuo: ", result)

	// Incremental
	x++
	fmt.Println("Incremental: ", x)

	// Decremental
	x--
	fmt.Println("Decremental: ", x)

}
