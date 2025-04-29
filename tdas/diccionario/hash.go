package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const TAMAÑO = 17

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}
type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{
		make([]TDALista.Lista[parClaveValor[K, V]], TAMAÑO),
		TAMAÑO,
		0,
	}
}

func (h hashAbierto[K, V]) Guardar(clave K, dato V) {
	//TODO implement me
	panic("implement me")
}

func (h hashAbierto[K, V]) Pertenece(clave K) bool {
	//TODO implement me
	panic("implement me")
}

func (h hashAbierto[K, V]) Obtener(clave K) V {
	//TODO implement me
	panic("implement me")
}

func (h hashAbierto[K, V]) Borrar(clave K) V {
	//TODO implement me
	panic("implement me")
}

func (h hashAbierto[K, V]) Cantidad() int {
	//TODO implement me
	panic("implement me")
}

func (h hashAbierto[K, V]) Iterar(f func(clave K, dato V) bool) {
	//TODO implement me
	panic("implement me")
}

func (h hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	//TODO implement me
	panic("implement me")
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func Hash[K comparable](clave K) uint32 {
	data := convertirABytes(clave)
	var hash uint32 = 0
	for _, b := range data {
		hash += uint32(b)
		hash += hash << 10
		hash ^= hash >> 6
	}
	hash += hash << 3
	hash ^= hash >> 11
	hash += hash << 15
	return hash
}

//primitivas del iterador
// hay que hacer el struct del iterador del hash
/*
func (i *) HaySiguiente() bool {

}

func (i *) VerActual() (K, V) {

}

func (i *) Siguiente() {

}
*/
