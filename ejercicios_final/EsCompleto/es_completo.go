package escompleto

import (
	"tdas/cola"
	"tdas/diccionario"
)

/*Implementar un algoritmo que, dado un árbol binario, determine si el mismo es completo (es decir, que todos los niveles
que tenga estén completos). Indicar y justificar la complejidad del algoritmo implementado.*/

type arbol[T any] struct {
izq *arbol[T]
der *arbol[T]
clave T
}

type nodoNivel[T any] struct {
	nodo *arbol[T]
	nivel int
}

func (ab *arbol[T]) EsCompleto() bool {
	dicc := diccionario.CrearHash[int, int]()
	c := cola.CrearColaEnlazada[nodoNivel[T]]()
	c.Encolar(nodoNivel[T]{ab, 0})
	for !c.EstaVacia() {
		n := c.Desencolar()
		if !dicc.Pertenece(n.nivel) {
			dicc.Guardar(n.nivel, 0)
		}
		dicc.Guardar(n.nivel, dicc.Obtener(n.nivel)+1)
		
		if n.nodo.izq != nil {
			c.Encolar(nodoNivel[T]{n.nodo.izq, n.nivel+1})
		}
		if n.nodo.der != nil {
			c.Encolar(nodoNivel[T]{n.nodo.der, n.nivel+1})
		}
	}
	for iter := dicc.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		nivel, cant_nodos := iter.VerActual()
		if 2^nivel != cant_nodos {
			return false
		}
	}
	return true
}

func (ab *arbol[T]) EsLleno() bool {
	arr := []*arbol[T]{}
	c := cola.CrearColaEnlazada[*arbol[T]]()
	c.Encolar(ab)
	for !c.EstaVacia() {
		n := c.Desencolar()
		arr = append(arr, n)
		if n.izq != nil {
			c.Encolar(n.izq)
		} else if n.izq == nil {
			c.Encolar(nil)
		}
		if n.der != nil {
			c.Encolar(n.der)
		} else if n.der == nil {
			c.Encolar(nil)
		}
	}
	hayNil := false
	for _, nodo := range arr {
		if nodo == nil {
			hayNil = true
		}
		if hayNil && nodo != nil {
			return false
		}
	}
	return true
}