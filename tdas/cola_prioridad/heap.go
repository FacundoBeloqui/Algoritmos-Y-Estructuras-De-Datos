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
	heap.cant++
	heap.datos[heap.cant] = elemento
	upheap(heap.datos, heap.cant, heap.cmp)
}

func (heap *colaConPrioridad[T]) VerMax() T {
	heap.verifcarColaVacia()
	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Desencolar() T {
	heap.verifcarColaVacia()
	elemento := heap.datos[0]



	return elemento
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func upheap[T any] (datos []T, i int, funcion_cmp func(T, T) int){
	for /*no se bien que poner aca*/ {
		padre := (i-1)/2
		if funcion_cmp(datos[i], datos[padre]) > 0{
			datos[i], datos[padre] = datos[padre], datos[i]
			i = padre
		} else {
			break
		}
	}

}

func downheap[T any] (datos[]T, i int, funcion_cmp func(T, T) int){


}

func HeapSort[T any](datos []T, funcion_cmp func(T, T) int){

}

func (heap *colaConPrioridad[T]) verifcarColaVacia(){
	if heap.EstaVacia(){
		panic("La cola esta vacia")
	}
}