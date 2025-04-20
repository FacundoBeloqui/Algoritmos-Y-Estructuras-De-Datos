package main

import (
	"bufio"
	"fmt"
	"os"
	"tdas/cola"
	"tdas/pila"
	"tp1/calculadoraPolaca"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var lineas []string
	for s.Scan() {
		lineas = append(lineas, s.Text())
	}

	for _, linea := range lineas {
		c := cola.CrearColaEnlazada[string]()
		p := pila.CrearPilaDinamica[calculadoraPolaca.Caracter]()
		tokens := calculadoraPolaca.SepararCadena(linea)
		for _, token := range tokens {
			calculadoraPolaca.ManejarToken(c, p, token)
		}
		calculadoraPolaca.VaciarPilaRestante(c, p)
		for !c.EstaVacia() {
			fmt.Printf("%s ", c.Desencolar())
		}
		fmt.Printf("\n")
	}
}
