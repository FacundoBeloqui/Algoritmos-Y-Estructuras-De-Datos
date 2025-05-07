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

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return abb.perteneceRec(abb.raiz, clave)
}

func (abb *abb[K, V]) perteneceRec(nodo *nodoAbb[K, V], clave K) bool {
	if nodo == nil {
		return false
	}
	if abb.cmp(clave, nodo.clave) < 0 {
		return abb.perteneceRec(nodo.izquierdo, clave)
	} else if abb.cmp(clave, nodo.clave) > 0 {
		return abb.perteneceRec(nodo.derecho, clave)
	}
	return true
}

func (abb *abb[K, V]) Obtener(clave K) V {
	return abb.obtenerRec(abb.raiz, clave)
}

func (abb *abb[K, V]) obtenerRec(nodo *nodoAbb[K, V], clave K) V {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	if abb.cmp(clave, nodo.clave) == 0 {
		return nodo.dato
	}
	if abb.cmp(clave, nodo.clave) < 0 {
		return abb.obtenerRec(nodo.izquierdo, clave)
	}
	return abb.obtenerRec(nodo.derecho, clave)
}

func (abb *abb[K, V]) Borrar(clave K) V {
	return abb.borrarRec(abb.raiz, clave)
}
func (abb *abb[K, V]) borrarRec(nodo *nodoAbb[K, V], clave K) V {
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	if abb.cmp(clave, nodo.clave) < 0 {
		return abb.borrarRec(nodo.izquierdo, clave)
	} else if abb.cmp(clave, nodo.clave) > 0 {
		return abb.borrarRec(nodo.derecho, clave)
	} else {
		if nodo.izquierdo == nil && nodo.derecho == nil {
			nodo = nil
		} else if nodo.izquierdo == nil && nodo.derecho != nil {
			nodo = nodo.derecho
		} else if nodo.izquierdo != nil && nodo.derecho == nil {
			nodo = nodo.izquierdo
		}
	}
	panic("ni idea que hacer pero va pora aca xd")
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
