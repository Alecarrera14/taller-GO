package main

import (
	"fmt"
	"funciones"
)

func main() {
	funciones.Saludar("Ana")

	saludo := funciones.Saludar2("Juan")
	fmt.Println(saludo)

	x := "Mundo"
	y := "Hola"
	x, y = funciones.Swap(x, y)
	fmt.Println(x, y)

	resultado, resto := funciones.Dividir(10, 3)
	fmt.Println(resultado, resto)
}
