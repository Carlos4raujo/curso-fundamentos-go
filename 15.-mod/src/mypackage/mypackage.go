package mypackage

import "fmt"

// Car mayuscula da acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// Car mayuscula da acceso privado
type carPrivate struct {
	brand string
	year  int
}

// PrintMessae imprime un mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}
