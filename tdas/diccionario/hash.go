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
		h.cantidad++
	}
	/*
		celda := Hash(clave) % TAMAÑO
		iterador := h.tabla[celda].Iterador()
		for iterador.HaySiguiente() {
			if iterador.VerActual().clave == clave {
				iterador.Borrar()
				h.cantidad--
				break
			}
			iterador.Siguiente()
		}
		h.tabla[celda].InsertarUltimo(ParClaveValor[K, V]{clave, dato})
		h.cantidad++
		}

	*/
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
	for _, lista := range h.tabla {
		if lista != nil {
			iterador := lista.Iterador()
			for iterador.HaySiguiente() {
				if !f(iterador.VerActual().clave, iterador.VerActual().dato) {
					break
				}
				iterador.Siguiente()
			}
		}
	}
}

type iterDiccionario[K comparable, V any] struct {
	hash          *hashAbierto[K, V]
	iteradorLista TDALista.IteradorLista[ParClaveValor[K, V]]
	posicion      int
}

func (h *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterDiccionario[K, V]{
		h,
		nil,
		0,
	}

	for iter.posicion < TAMAÑO {
		lista := iter.hash.tabla[iter.posicion]
		if !lista.EstaVacia() {
			iter.iteradorLista = lista.Iterador()
			break
		}
		iter.posicion++
	}
	return iter
}

func (i *iterDiccionario[K, V]) HaySiguiente() bool {
	for i.posicion < TAMAÑO {
		lista := i.hash.tabla[i.posicion]
		if !lista.EstaVacia() && lista.Iterador().HaySiguiente() {
			return true
		}
		i.posicion++
	}
	return false
}

func (i *iterDiccionario[K, V]) VerActual() (K, V) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return i.iteradorLista.VerActual().clave, i.iteradorLista.VerActual().dato
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
