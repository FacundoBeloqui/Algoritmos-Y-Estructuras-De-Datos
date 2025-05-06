package diccionario

import (
	"fmt"
	"hash/fnv"
	TDALista "tdas/lista"
)

const (
	TAMAÑO_INICIAL     = 17
	FACTOR_CARGA_MAX   = 0.75
	FACTOR_CARGA_MIN   = 0.25
	FACTOR_REDIMENSION = 2
)

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
	tabla := make([]TDALista.Lista[parClaveValor[K, V]], TAMAÑO_INICIAL)
	for i := range tabla {
		tabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
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

func hashing[K comparable](clave K) uint32 {
	h := fnv.New32a()
	h.Write(convertirABytes(clave))
	return h.Sum32()
}

func (hash *hashAbierto[K, V]) Guardar(clave K, dato V) {
	if float32(hash.cantidad)/float32(hash.tam) > FACTOR_CARGA_MAX {
		hash.redimensionar(hash.tam * FACTOR_REDIMENSION)
	}
	iterador := hash.encontrarCampo(clave)
	if iterador != nil {
		iterador.Borrar()
		iterador.Insertar(parClaveValor[K, V]{clave, dato})
		return
	}
	celda := hashing(clave) % uint32(hash.tam)
	hash.tabla[celda].InsertarUltimo(parClaveValor[K, V]{clave, dato})
	hash.cantidad++
}

func (hash *hashAbierto[K, V]) Pertenece(clave K) bool {
	return hash.encontrarCampo(clave) != nil
}

func (hash *hashAbierto[K, V]) Obtener(clave K) V {
	iterador := hash.encontrarCampo(clave)
	hash.verificarIterador(iterador)
	return iterador.VerActual().dato
}

func (hash *hashAbierto[K, V]) Borrar(clave K) V {
	iterador := hash.encontrarCampo(clave)
	hash.verificarIterador(iterador)
	campo := iterador.Borrar()
	hash.cantidad--
	if hash.tam > TAMAÑO_INICIAL && float32(hash.cantidad)/float32(hash.tam) < FACTOR_CARGA_MIN {
		hash.redimensionar(hash.tam / FACTOR_REDIMENSION)
	}
	return campo.dato
}

func (hash *hashAbierto[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashAbierto[K, V]) encontrarCampo(clave K) TDALista.IteradorLista[parClaveValor[K, V]] {
	celda := hashing(clave) % uint32(hash.tam)
	iterador := hash.tabla[celda].Iterador()
	for iterador.HaySiguiente() {
		if iterador.VerActual().clave == clave {
			return iterador
		}
		iterador.Siguiente()
	}
	return nil
}

func (hash *hashAbierto[K, V]) redimensionar(nuevoTam int) {
	nuevaTabla := make([]TDALista.Lista[parClaveValor[K, V]], nuevoTam)
	for i := range nuevaTabla {
		nuevaTabla[i] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	for _, lista := range hash.tabla {
		iterador := lista.Iterador()
		for iterador.HaySiguiente() {
			nuevaCelda := hashing(iterador.VerActual().clave) % uint32(nuevoTam)
			nuevaTabla[nuevaCelda].InsertarUltimo(iterador.VerActual())
			iterador.Siguiente()
		}
	}
	hash.tabla = nuevaTabla
	hash.tam = nuevoTam
}

func (hash *hashAbierto[K, V]) verificarIterador(iterador TDALista.IteradorLista[parClaveValor[K, V]]) {
	if iterador == nil {
		panic("La clave no pertenece al diccionario")
	}
}

func (hash *hashAbierto[K, V]) Iterar(f func(clave K, dato V) bool) {
	for _, lista := range hash.tabla {
		iterador := lista.Iterador()
		for iterador.HaySiguiente() {
			if !f(iterador.VerActual().clave, iterador.VerActual().dato) {
				return
			}
			iterador.Siguiente()
		}
	}
}

type iterDiccionario[K comparable, V any] struct {
	hash          *hashAbierto[K, V]
	iteradorLista TDALista.IteradorLista[parClaveValor[K, V]]
	posicion      int
}

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterDiccionario[K, V]{
		hash,
		nil,
		-1,
	}
	posInicial := iter.encontrarLista(0)
	if posInicial != -1 {
		iter.posicion = posInicial
	} else {
		iter.posicion = 0
	}
	iter.iteradorLista = iter.hash.tabla[iter.posicion].Iterador()
	return iter
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	return iter.iteradorLista.HaySiguiente()
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	iter.verificarIterador()
	return iter.iteradorLista.VerActual().clave, iter.iteradorLista.VerActual().dato
}

func (iter *iterDiccionario[K, V]) Siguiente() {
	iter.verificarIterador()
	iter.iteradorLista.Siguiente()
	if !iter.iteradorLista.HaySiguiente() {
		nuevaPos := iter.encontrarLista(iter.posicion + 1)
		if nuevaPos == -1 {
			return
		}
		iter.posicion = nuevaPos
		iter.iteradorLista = iter.hash.tabla[iter.posicion].Iterador()
	}
}

func (iter *iterDiccionario[K, V]) encontrarLista(desde int) int {
	for i := desde; i < iter.hash.tam; i++ {
		if !iter.hash.tabla[i].EstaVacia() {
			return i
		}
	}
	return -1
}

func (iter *iterDiccionario[K, V]) verificarIterador() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
