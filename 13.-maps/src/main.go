package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// Recorrer map
	for i, valor := range m {
		fmt.Println(i, valor)
	}

	// Enocontrar un valor
	value := m["Jose"]
	fmt.Println(value)

	// Llaves inexistentes
	value1, ok := m["Joseph"]
	fmt.Println(value1, ok)
}
