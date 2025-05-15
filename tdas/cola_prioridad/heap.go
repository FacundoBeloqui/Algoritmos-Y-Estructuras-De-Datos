package cola_prioridad

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: make([]T, 0),
		cant:  0,
		cmp:   funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: arreglo,
		cant:  len(arreglo),
		cmp:   funcion_cmp,
	}
}
func (heap *colaConPrioridad[T]) calcularPosicionHijoIzquierdo(posicion int) int {
	return posicion*2 + 1
}
func (heap *colaConPrioridad[T]) calcularPosicionHijoDerecho(posicion int) int {
	return posicion*2 + 2
}
func (heap *colaConPrioridad[T]) calcularPosicionPadre(posicion int) int {
	return (posicion - 1) / 2
}
func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *colaConPrioridad[T]) Encolar(elemento T) {
	heap.datos = append(heap.datos, elemento)
	heap.cant++
	heap.upheap()
}
func (heap *colaConPrioridad[T]) upheap() {
	posicionHijo := heap.cant
	posicionPadre := heap.calcularPosicionPadre(posicionHijo)
	padre := heap.datos[posicionPadre]
	hijo := heap.datos[posicionHijo]
	for heap.cmp(hijo, padre) > 0 || posicionHijo != 0 {
		posicionHijo = posicionPadre
		posicionPadre = heap.calcularPosicionPadre(posicionHijo)
	}
}
func (heap *colaConPrioridad[T]) VerMax() T {
	heap.verifcarColaVacia()
	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Desencolar() T {
	heap.verifcarColaVacia()
	return heap.datos[len(heap.datos)-1]
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {

}

func (heap *colaConPrioridad[T]) verifcarColaVacia() {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
}
