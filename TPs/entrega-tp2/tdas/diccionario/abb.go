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
	nodo, padre := abb.buscarConPadre(abb.raiz, clave, nil)
	if nodo != nil {
		nodo.dato = dato
		return
	}
	abb.cantidad++
	if padre == nil {
		abb.raiz = crearNodo(clave, dato)
	} else if abb.cmp(clave, padre.clave) < 0 {
		padre.izquierdo = crearNodo(clave, dato)
	} else {
		padre.derecho = crearNodo(clave, dato)
	}
}

func (abb *abb[K, V]) buscarConPadre(nodo *nodoAbb[K, V], clave K, padre *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo == nil {
		return nil, padre
	}

	cmp := abb.cmp(clave, nodo.clave)
	if cmp == 0 {
		return nodo, padre
	} else if cmp < 0 {
		return abb.buscarConPadre(nodo.izquierdo, clave, nodo)
	} else {
		return abb.buscarConPadre(nodo.derecho, clave, nodo)
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := abb.buscarConPadre(abb.raiz, clave, nil)
	return nodo != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	nodo, _ := abb.buscarConPadre(abb.raiz, clave, nil)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo, padre := abb.buscarConPadre(abb.raiz, clave, nil)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	valor := nodo.dato
	abb.cantidad--
	var reemplazo *nodoAbb[K, V]
	if nodo.izquierdo == nil && nodo.derecho == nil {
		reemplazo = nil
	} else if nodo.izquierdo == nil || nodo.derecho == nil {
		if nodo.izquierdo != nil {
			reemplazo = nodo.izquierdo
		} else {
			reemplazo = nodo.derecho
		}
	} else {
		siguiente := abb.buscarSiguiente(nodo.derecho)
		abb.cantidad++
		abb.Borrar(siguiente.clave)
		nodo.clave = siguiente.clave
		nodo.dato = siguiente.dato
		return valor
	}
	if padre == nil {
		abb.raiz = reemplazo
	} else if padre.izquierdo == nodo {
		padre.izquierdo = reemplazo
	} else {
		padre.derecho = reemplazo
	}
	return valor
}

func (abb *abb[K, V]) buscarSiguiente(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	if nodo == nil || nodo.izquierdo == nil {
		return nodo
	}
	return abb.buscarSiguiente(nodo.izquierdo)
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

//<----------- ITERADOR INTERNO ----------->

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	abb.IterarRango(nil, nil, f)
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.raiz.iterarRango(desde, hasta, visitar, abb.cmp)
}

func (nodo *nodoAbb[K, V]) iterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool, funcCmp func(K, K) int) bool {
	if nodo == nil {
		return true
	}

	if desde == nil || funcCmp(nodo.clave, *desde) > 0 {
		if !nodo.izquierdo.iterarRango(desde, hasta, visitar, funcCmp) {
			return false
		}
	}

	if (desde == nil || funcCmp(nodo.clave, *desde) >= 0) && (hasta == nil || funcCmp(nodo.clave, *hasta) <= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}

	if hasta == nil || funcCmp(nodo.clave, *hasta) < 0 {
		if !nodo.derecho.iterarRango(desde, hasta, visitar, funcCmp) {
			return false
		}
	}
	return true
}

//<----------- ITERADOR EXTERNO ----------->

type iterAbb[K comparable, V any] struct {
	pila  TDAPila.Pila[nodoAbb[K, V]]
	cmp   funcCmp[K]
	desde *K
	hasta *K
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := &iterAbb[K, V]{
		pila:  TDAPila.CrearPilaDinamica[nodoAbb[K, V]](),
		cmp:   abb.cmp,
		desde: desde,
		hasta: hasta,
	}
	iter.apilarDesdeHasta(abb.raiz)
	return iter
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.Desapilar()
	iter.apilarDesdeHasta(nodo.derecho)
}

func (iter *iterAbb[K, V]) apilarDesdeHasta(nodo *nodoAbb[K, V]) {
	for nodo != nil {
		if iter.desde != nil && iter.cmp(nodo.clave, *iter.desde) < 0 {
			nodo = nodo.derecho
		} else if iter.hasta != nil && iter.cmp(nodo.clave, *iter.hasta) > 0 {
			nodo = nodo.izquierdo
		} else {
			iter.pila.Apilar(*nodo)
			nodo = nodo.izquierdo
		}
	}
}
