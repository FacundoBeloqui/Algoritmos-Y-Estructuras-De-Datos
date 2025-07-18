package main

import (
	"ejercicios_final/iterador_interno"
	"fmt"
	"tdas/lista"
)
func duplicar(n int) int {
	return n*2
}
func main() {
	lista := lista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(7)
	lista.InsertarUltimo(12)
	nuevaLista := iterador_interno.Map[int](lista, duplicar)
	nuevaLista.Iterar(func(i int) bool {
		fmt.Println(i)
		return true
	})
}