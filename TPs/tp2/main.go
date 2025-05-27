package main

import (
	"fmt"
	"tp2/algueiza"
)

func main() {
	vuelos := algueiza.AgregarArchivo("vuelos-algueiza/vuelos-algueiza-parte-01.csv")
	for iter := vuelos.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		fmt.Printf("%d, %v \n", clave, valor)
	}
}
