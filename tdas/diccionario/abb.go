package diccionario

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

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	if abb.raiz == nil {
		abb.raiz = &nodoAbb[K, V]{
			nil,
			nil,
			clave,
			dato,
		}
	} else {
		abb.guardarRec(abb.raiz, clave, dato)
	}
}

func (abb *abb[K, V]) guardarRec(nodo *nodoAbb[K, V], clave K, dato V) *nodoAbb[K, V] {
	if nodo == nil {
		abb.cantidad++
		return &nodoAbb[K, V]{nil, nil, clave, dato}
	}
	if abb.cmp(clave, nodo.clave) < 0 {
		return abb.guardarRec(nodo.izquierdo, clave, dato)
	} else if abb.cmp(clave, nodo.clave) > 0 {
		return abb.guardarRec(nodo.derecho, clave, dato)
	}
	return &nodoAbb[K, V]{nodo.izquierdo, nodo.derecho, clave, dato}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return abb.perteneceRec(abb.raiz, clave)
}

func (abb *abb[K, V]) perteneceRec(nodo *nodoAbb[K, V], clave K) bool {
	if abb.cmp(clave, nodo.clave) == 0 {
		return true
	}
	if abb.cmp(clave, nodo.clave) < 0 {
		return abb.perteneceRec(nodo.izquierdo, clave)
	} else if abb.cmp(clave, nodo.clave) > 0 {
		return abb.perteneceRec(nodo.derecho, clave)
	}
	return false
}

func (abb *abb[K, V]) Obtener(clave K) V {
	return abb.obtenerRec(abb.raiz, clave)
}

func (abb *abb[K, V]) obtenerRec(nodo *nodoAbb[K, V], clave K) V {
	if abb.cmp(clave, nodo.clave) == 0 {
		return nodo.dato
	}
	if abb.cmp(clave, nodo.clave) < 0 {
		return abb.obtenerRec(nodo.izquierdo, clave)
	} else if abb.cmp(clave, nodo.clave) > 0 {
		return abb.obtenerRec(nodo.derecho, clave)
	}
	panic("La clave no pertenece al diccionario")
}

func (abb *abb[K, V]) Borrar(clave K) V {
	//TODO implement me
	panic("implement me")
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	//TODO implement me
	panic("implement me")
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	//TODO implement me
	panic("implement me")
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	//TODO implement me
	panic("implement me")
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	//TODO implement me
	panic("implement me")
}
