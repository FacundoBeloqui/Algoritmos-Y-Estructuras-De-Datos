package main

import (
	"fmt"
	"strconv"
)

type Palo int

const (
	PICAS Palo = iota
	CORAZONES
	TREBOLES
	DIAMANTES
)

type Carta struct {
	numero string
	palo   Palo
}

// numeroCartaAIndice pasa del numero de la carta (que podría ser 'A', 'J', 'Q', 'K') al indice que corresponde
func numeroCartaAIndice(carta Carta) int {
	switch carta.numero {
	case "A":
		return 0
	case "K":
		return 12
	case "Q":
		return 11
	case "J":
		return 10
	default:
		v, _ := strconv.Atoi(carta.numero)
		return v - 1
	}
}
func paloAIndice(carta Carta) int {
	switch carta.palo {
	case PICAS:
		return 0
	case CORAZONES:
		return 1
	case TREBOLES:
		return 2
	case DIAMANTES:
		return 3
	default:
		return -1
	}
}
func countingSort(cartas []Carta, k int, criterio func(Carta)int) []Carta {
	frecuencias := make([]int, k)
	for _, carta := range cartas {
		frecuencias[criterio(carta)]++
	}
	sumasAcumuladas := make([]int, k)
	for i := 1; i < k; i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}
	ordenadas := make([]Carta, len(cartas))
	for _, carta := range cartas {
		indice := criterio(carta)
		ordenadas[sumasAcumuladas[indice]] = carta
		sumasAcumuladas[indice]++
	}
	return ordenadas
}

func ordenarPorNumero(cartas []Carta) []Carta {
	return countingSort(cartas, 13, numeroCartaAIndice)
}

func ordenarPorPalo(cartas []Carta) []Carta {
	return countingSort(cartas, 4, paloAIndice)
}

func ordenarCartasRadix(cartas []Carta) []Carta {
	return ordenarPorPalo(ordenarPorNumero(cartas))
}

func main() {
	mazo := []Carta{
		{"8", PICAS},
		{"7", TREBOLES},
		{"3", CORAZONES},
		{"8", TREBOLES},
		{"7", PICAS},
		{"A", CORAZONES},
		{"7", TREBOLES},
		{"3", DIAMANTES},
		{"Q", PICAS},
		{"2", DIAMANTES},
		{"A", DIAMANTES},
		{"2", CORAZONES},
		{"8", CORAZONES},
		{"J", DIAMANTES},
	}
	fmt.Println(ordenarCartasRadix(mazo))
}
