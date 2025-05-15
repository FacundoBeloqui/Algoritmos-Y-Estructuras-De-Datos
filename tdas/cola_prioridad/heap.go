package cola_prioridad

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T]{
	return &colaConPrioridad[T]{
		datos: make([]T, 0),
		cant: 0,
		cmp: funcion_cmp,
	}
}



func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T]{
	return &colaConPrioridad[T]{
		datos: arreglo,
		cant: len(arreglo),
		cmp: funcion_cmp,
	}
}

func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *colaConPrioridad[T]) Encolar(elemento T) {

}

func (heap *colaConPrioridad[T]) VerMax() T {
	heap.verifcarColaVacia()
	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Desencolar() T {
	heap.verifcarColaVacia()
	
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int){

}

func (heap *colaConPrioridad[T]) verifcarColaVacia(){
	if heap.EstaVacia(){
		panic("La cola esta vacia")
	}
}