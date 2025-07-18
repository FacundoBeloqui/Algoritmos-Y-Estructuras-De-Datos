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

func ordenarCartasCountingCompleto(cartas []Carta) []Carta {
	frequencias := make([]int, 13)
	for _, carta := range cartas {
		frequencias[numeroCartaAIndice(carta)]++
	}

	sumas_acumuladas := make([]int, 13)
	for i := 1; i < 13; i++ {
		sumas_acumuladas[i] = sumas_acumuladas[i-1] + frequencias[i-1]
	}

	ordenadas := make([]Carta, len(cartas))
	for _, carta := range cartas {
		indice := numeroCartaAIndice(carta)
		ordenadas[sumas_acumuladas[indice]] = carta
		sumas_acumuladas[indice]++
	}
	return ordenadas
}

/*func ordenarCartasCountingSimplificado(cartas []Carta) []Carta {
	colas := make([]Cola, 13)
	for i := range colas {
		colas[i] = TDACola.CrearColaEnlazada()
	}
	for _, carta := range cartas {
		colas[numeroCartaAIndice(carta)].Encolar(carta)
	}

	ordenadas := make([]Carta, len(cartas))
	indice := 0
	for _, cola := range colas {
		for !cola.EstaVacia() {
			ordenadas[indice] = cola.Desencolar()
			indice++
		}
	}
	return ordenadas
}*/

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
	fmt.Println(ordenarCartasCountingCompleto(mazo))
}
