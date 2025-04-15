package main

import (
	"Guia/ejerciciosTDAs"
	"fmt"
	TDAPila "tdas/pila"
)

func main() {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(4)
	pila.Apilar(3)
	pila.Apilar(15)
	pila.Apilar(6)
	pila.Apilar(0)
	pila.Apilar(8)
	pila.Apilar(9)
	pila.Apilar(20)
	cola := ejerciciosTDAs.Distribuir(pila)
	cola.Desencolar()
	fmt.Printf("%d", cola.VerPrimero())
}
