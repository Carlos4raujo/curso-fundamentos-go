package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram *= 2
}

func main() {
	a := 50
	b := &a

	fmt.Println(a)
	fmt.Println(b)  // -> '&' es para acceder a la dirección de memoria
	fmt.Println(*b) // -> '*' es para acceder al valor de la dirección en memoria

	*b = 100
	fmt.Println(a)

	myPC := pc{ram: 16, disk: 200, brand: "msi"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)
	myPC.duplicateRAM()

	fmt.Println(myPC)

}
