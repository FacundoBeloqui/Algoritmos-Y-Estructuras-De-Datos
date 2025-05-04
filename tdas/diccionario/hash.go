package diccionario

import (
	"fmt"
	"hash/fnv"
	TDALista "tdas/lista"
)

const TAMAÑO_INICIAL = 17

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
	tabla := make([]TDALista.Lista[ParClaveValor[K, V]], TAMAÑO_INICIAL)
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[ParClaveValor[K, V]]()
	}
	return &hashAbierto[K, V]{
		tabla,
		TAMAÑO_INICIAL,
		0,
	}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

/*func Hash[K comparable](clave K) uint32 {
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
}*/

func Hash[K comparable](clave K) uint32 {
	h := fnv.New32a()
	h.Write(convertirABytes(clave))
	return h.Sum32()
}


func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {	
	celda := Hash(clave) % uint32(h.tam)
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

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {
	celda := Hash(clave) % uint32(h.tam)
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
	iterador := h.encontrarIterador(clave)
	return iterador.VerActual().dato
}

func (h *hashAbierto[K, V]) Borrar(clave K) V {
	iterador := h.encontrarIterador(clave)
	dato := iterador.VerActual().dato
	iterador.Borrar()
	h.cantidad--
	return dato
}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) encontrarIterador(clave K) TDALista.IteradorLista[ParClaveValor[K, V]] {
	celda := Hash(clave) % uint32(h.tam)
	iterador := h.tabla[celda].Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual().clave == clave {
			return iterador
		}
		iterador.Siguiente()
	}
	panic("La clave no pertenece al diccionario")
}

func (h *hashAbierto[K, V]) Iterar(f func(clave K, dato V) bool) {
	for _, lista := range h.tabla {
		if lista != nil {
			iterador := lista.Iterador()
			for iterador.HaySiguiente() {
				if !f(iterador.VerActual().clave, iterador.VerActual().dato) {
					return
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

	iter.encontrarLista()
	return iter
}

func (i *iterDiccionario[K, V]) HaySiguiente() bool {
	return i.iteradorLista != nil
}


func (i *iterDiccionario[K, V]) VerActual() (K, V) {
	i.verificarIterador()
	return i.iteradorLista.VerActual().clave, i.iteradorLista.VerActual().dato
}

func (i *iterDiccionario[K, V]) Siguiente() {
	i.verificarIterador()
	if i.iteradorLista.HaySiguiente() {
		i.iteradorLista.Siguiente()
	} else {
		i.posicion++
		i.encontrarLista()
	}
}

func (i *iterDiccionario[K, V]) encontrarLista(){
	for i.posicion < i.hash.tam{
		lista := i.hash.tabla[i.posicion]
		if !lista.EstaVacia() {
			i.iteradorLista = lista.Iterador()
			i.iteradorLista.Siguiente()
			return
		}
		i.posicion++
	}
	i.iteradorLista = nil
}

func (i *iterDiccionario[K, V]) verificarIterador(){
	if i.iteradorLista == nil {
		panic("El iterador termino de iterar")
	}
}