package main

import (
	"fmt"
	"tdas/pila"
)

func main() {
	pila1 := pila.CrearPilaDinamica[string]()
	fmt.Println("mensaje:", pila1.Desapilar())
}
