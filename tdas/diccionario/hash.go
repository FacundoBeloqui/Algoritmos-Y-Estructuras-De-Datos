package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const TAMAÑO = 17

type ParClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}
type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[ParClaveValor[K, V]]
	tam      int
	cantidad int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tabla := make([]TDALista.Lista[ParClaveValor[K, V]], TAMAÑO)
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[ParClaveValor[K, V]]()
	}
	return &hashAbierto[K, V]{
		tabla,
		TAMAÑO,
		0,
	}
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

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {
	celda := Hash(clave) % TAMAÑO
	iterador := h.tabla[celda].Iterador()
	if h.Pertenece(clave) {
		for iterador.HaySiguiente() {
			if iterador.VerActual().clave == clave {
				iterador.Borrar()
				iterador.Insertar(ParClaveValor[K, V]{clave, dato})
			}
			iterador.Siguiente()
		}
	} else {
		h.tabla[celda].InsertarUltimo(ParClaveValor[K, V]{clave, dato})
	}
	h.cantidad++

}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	celda := Hash(clave) % TAMAÑO
	iterador := h.tabla[celda].Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual().clave == clave {
			return true
		}
		iterador.Siguiente()
	}
	return false
}

func (h *hashAbierto[K, V]) Obtener(clave K) V {
	celda := Hash(clave) % TAMAÑO
	iterador := h.tabla[celda].Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual().clave == clave {
			return iterador.VerActual().dato
		}
		iterador.Siguiente()
	}
	panic("La clave no pertenece al diccionario")
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	celda := Hash(clave) % TAMAÑO
	iterador := h.tabla[celda].Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual().clave == clave {
			dato := iterador.VerActual().dato
			iterador.Borrar()
			h.cantidad--
			return dato
		}
		iterador.Siguiente()
	}

	panic("La clave no pertenece al diccionario")
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Iterar(f func(clave K, dato V) bool) {
	//TODO implement me
	panic("implement me")
}

func (i *iterDiccionario[K, V]) Siguiente() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	if i.iteradorLista.HaySiguiente() {
		i.iteradorLista.Siguiente()
	} else {
		i.posicion++
	}
}
