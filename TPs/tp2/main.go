package main

import (
	"fmt"
	"strings"
	"tdas/diccionario"
	"tp2/algueiza"
)

func main() {
	diccFechas := diccionario.CrearABB[string, algueiza.VueloImpl](strings.Compare)
	diccNumerosVuelo := diccionario.CrearHash[int, string]()
	algueiza.AgregarArchivo("vuelos-algueiza/vuelos-algueiza-parte-01.csv", diccFechas, diccNumerosVuelo)
	algueiza.AgregarArchivo("vuelos-algueiza/vuelos-algueiza-parte-02.csv", diccFechas, diccNumerosVuelo)
	algueiza.AgregarArchivo("vuelos-algueiza/vuelos-algueiza-parte-03.csv", diccFechas, diccNumerosVuelo)
	for iter := diccFechas.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		fmt.Printf("%s, %v \n", clave, valor)
	}
}
