package main

import (
	"fmt"
	"strings"
	"tdas/diccionario"
	"tp2/algueiza"
)

func main() {
	diccFechas := diccionario.CrearABB[string, algueiza.VueloImpl](strings.Compare)
	diccNumerosVuelo := diccionario.CrearHash[int, algueiza.VueloImpl]()
	//algueiza.AgregarArchivo("vuelos-algueiza/vuelos-algueiza-parte-01.csv", diccFechas, diccNumerosVuelo)
	algueiza.AgregarArchivo("vuelos-algueiza/vuelos-algueiza-parte-02.csv", diccFechas, diccNumerosVuelo)
	algueiza.AgregarArchivo("vuelos-algueiza/vuelos-algueiza-parte-03.csv", diccFechas, diccNumerosVuelo)
	algueiza.AgregarArchivo("vuelos-algueiza/vuelos-algueiza-parte-04.csv", diccFechas, diccNumerosVuelo)
	algueiza.AgregarArchivo("vuelos-algueiza/vuelos-algueiza-parte-05.csv", diccFechas, diccNumerosVuelo)
	for iter := diccFechas.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		fmt.Printf("%s, %v \n", clave, valor)
	}
	algueiza.VerTablero(8, "asc", "2018-04-08T10:00:00", "2019-04-21T10:12:00", diccFechas)
	algueiza.InfoVuelo(1070, diccNumerosVuelo)
	algueiza.PrioridadVuelos(5, diccNumerosVuelo)
}
