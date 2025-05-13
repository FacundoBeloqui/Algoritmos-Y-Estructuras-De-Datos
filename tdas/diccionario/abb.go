package diccionario

import TDAPila "tdas/pila"

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{
		nil,
		0,
		funcion_cmp,
	}
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{
		nil,
		nil,
		clave,
		dato,
	}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	abb.raiz = abb.guardarRec(abb.raiz, clave, dato)
}

func (abb *abb[K, V]) guardarRec(nodo *nodoAbb[K, V], clave K, dato V) *nodoAbb[K, V] {
	if nodo == nil {
		abb.cantidad++
		return crearNodo(clave, dato)
	}
	if abb.cmp(clave, nodo.clave) < 0 {
		nodo.izquierdo = abb.guardarRec(nodo.izquierdo, clave, dato)
	} else if abb.cmp(clave, nodo.clave) > 0 {
		nodo.derecho = abb.guardarRec(nodo.derecho, clave, dato)
	} else {
		nodo.dato = dato
	}
	return nodo
}

func (abb *abb[K, V]) buscarNodo(nodo *nodoAbb[K, V], clave K) *nodoAbb[K, V] {
	if nodo == nil {
		return nil
	}
	if abb.cmp(clave, nodo.clave) == 0 {
		return nodo
	} else if abb.cmp(clave, nodo.clave) < 0 {
		return abb.buscarNodo(nodo.izquierdo, clave)
	}
	return abb.buscarNodo(nodo.derecho, clave)
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return abb.buscarNodo(abb.raiz, clave) != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo := abb.buscarNodo(abb.raiz, clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	var valor V
	abb.raiz, valor = abb.borrarRec(abb.raiz, clave)
	return valor
}
func (abb *abb[K, V]) borrarRec(nodo *nodoAbb[K, V], clave K) (*nodoAbb[K, V], V) {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}

	var valor V
	if abb.cmp(clave, nodo.clave) < 0 {
		nodo.izquierdo, valor = abb.borrarRec(nodo.izquierdo, clave)
	} else if abb.cmp(clave, nodo.clave) > 0 {
		nodo.derecho, valor = abb.borrarRec(nodo.derecho, clave)
	} else {
		valor = nodo.dato

		if nodo.izquierdo == nil && nodo.derecho == nil {
			abb.cantidad--
			return nil, valor
		} else if nodo.izquierdo == nil && nodo.derecho != nil {
			abb.cantidad--
			return nodo.derecho, valor
		} else if nodo.izquierdo != nil && nodo.derecho == nil {
			abb.cantidad--
			return nodo.izquierdo, valor
		}

		siguienteInorder := abb.buscarSiguiente(nodo.derecho)
		nodo.clave = siguienteInorder.clave
		nodo.dato = siguienteInorder.dato
		nodo.derecho, _ = abb.borrarRec(nodo.derecho, siguienteInorder.clave)
	}
	return nodo, valor
}

func (abb *abb[K, V]) buscarSiguiente(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	for nodo.izquierdo != nil {
		nodo = nodo.izquierdo
	}
	return nodo
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

//<----------- ITERADOR INTERNO ----------->

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	abb.IterarRango(nil, nil, f)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if abb.raiz == nil {
		return
	}
	abb.raiz.iterarRango(desde, hasta, visitar, abb.cmp)
}

func (nodo *nodoAbb[K, V]) iterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool, funcCmp func(K, K) int) {
	if nodo == nil {
		return
	}

	if desde == nil || funcCmp(nodo.clave, *desde) > 0 {
		nodo.izquierdo.iterarRango(desde, hasta, visitar, funcCmp)
	}

	if (desde == nil || funcCmp(nodo.clave, *desde) >= 0) && (hasta == nil || funcCmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return
		}
	}

	if hasta == nil || funcCmp(nodo.clave, *hasta) < 0 {
		nodo.derecho.iterarRango(desde, hasta, visitar, funcCmp)
	}
}

//<----------- ITERADOR EXTERNO ----------->

type iterAbb[K comparable, V any] struct {
	pila TDAPila.Pila[nodoAbb[K, V]]
	abb  abb[K, V]
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iterAbb[K, V]{
		pila: TDAPila.CrearPilaDinamica[nodoAbb[K, V]](),
		abb:  *abb,
	}
	iter.apilarHijosIzquierdos(abb.raiz)
	return iter
}
func (iter *iterAbb[K, V]) apilarHijosIzquierdos(nodo *nodoAbb[K, V]) {
	if nodo == nil {
		return
	}
	iter.pila.Apilar(*nodo)
	iter.apilarHijosIzquierdos(nodo.izquierdo)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iterAbb[K, V]{
		pila: TDAPila.CrearPilaDinamica[nodoAbb[K, V]](),
		abb:  *abb,
	}

	padre := abb.raiz.BuscarPadre(*desde, abb.cmp)
	iter.pila.Apilar(*padre)
	
	return iter

}

func (nodo *nodoAbb[K, V]) BuscarPadre(desde K, funcCmp func(K, K) int) *nodoAbb[K, V] {
	if nodo == nil {
		return nil
	}
	if funcCmp(desde, nodo.clave) == 0 {
		return nil
	}
	if funcCmp(desde, nodo.clave) < 0 {
		if nodo.izquierdo != nil && funcCmp(nodo.izquierdo.clave, desde) <= 0 && funcCmp(desde, nodo.izquierdo.clave) == 0 {
			return nodo
		}
		return nodo.izquierdo.BuscarPadre(desde, funcCmp)
	} else {
		if nodo.derecho != nil && funcCmp(nodo.derecho.clave, desde) >= 0 && funcCmp(desde, nodo.derecho.clave) == 0 {
			return nodo
		}
		return nodo.derecho.BuscarPadre(desde, funcCmp)
	}
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if iter.pila.EstaVacia() {
		panic("El iterador termino de iterar")
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	if iter.pila.EstaVacia() {
		panic("El iterador termino de iterar")
	}

	nodo := iter.pila.Desapilar()
	if nodo.derecho != nil {
		iter.pila.Apilar(*nodo.derecho)
		iter.apilarHijosIzquierdos(nodo.derecho)
	}
}