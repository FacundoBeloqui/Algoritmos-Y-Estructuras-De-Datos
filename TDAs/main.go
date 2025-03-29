package main

import (
	"TDAs/pila"
	"fmt"
)

func main() {
	pila1 := pila.CrearPilaDinamica[string]()
	fmt.Println("mensaje:", pila1.Desapilar())
}
