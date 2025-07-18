package main

import "fmt"

/*Realizar el seguimiento de aplicar CountingSort al siguiente conjunto de selecciones de fútbol, ordenando por la cantidad de copas
continentales que tienen (entre paréntesis se indica la cantidad en cada caso). Implementar dicho algoritmo, e indicar y justificar la
complejidad del mismo.
España (4) - Inglaterra (0) - Uruguay (15) - Perú (2) - Francia (2) -
Portugal (1) - Argentina (16) - Alemania (3) - Brasil (9) - Dinamarca (1) -
Nigeria (3) - Chequia (1) - Corea del Sur (2) - Egipto (7) - Ghana (4) - Japón (4)*/

type seleccion struct {
	nombre string
	cantidadCopas int
}
func cantidadCopasIndice(sele seleccion) int {
	return sele.cantidadCopas
}
func CountingSort(sele []seleccion) []seleccion {
	frecuencias := make([]int, 17)
	for _, s := range sele {
		frecuencias[cantidadCopasIndice(s)]++
	}
	sumasAcumuladas := make([]int, 17) 
	for i := 1; i < 17; i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i -1]
	}
	ordenados := make([]seleccion, len(sele))
	for _, s := range sele {
		indice := cantidadCopasIndice(s)
		ordenados[sumasAcumuladas[indice]] = s
		sumasAcumuladas[indice]++
	}
	return ordenados
}

func main() {
	selecciones := []seleccion {seleccion {"España", 4}, seleccion {"Inglaterra", 0}, seleccion {"Uruguay", 15}, seleccion {"Peru", 2},
	seleccion {"Francia", 2}, seleccion {"Portugal", 1}, seleccion {"Argentina", 16}, seleccion {"Alemania", 3}, seleccion {"España", 4}, 
	seleccion {"Brazil", 9}, seleccion {"España", 4}, seleccion {"Dinamarca", 1}, seleccion {"Nigeria", 3}, seleccion {"Chequia", 1},
	seleccion {"Corea Del Sur", 2}, seleccion {"Egipto", 7}, seleccion {"Ghana", 4}, seleccion {"Japon", 4}}
	fmt.Println(CountingSort(selecciones))
}