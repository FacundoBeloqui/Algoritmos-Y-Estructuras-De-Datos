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
	heap := &colaConPrioridad[T]{
		datos: arreglo,
		cant:  len(arreglo),
		cmp:   funcion_cmp,
	}
	Heapify(arreglo, len(arreglo), funcion_cmp)
	return heap
}

func calcularPosicionHijoIzquierdo(posicion int) int {
	return posicion*2 + 1
}
func calcularPosicionHijoDerecho(posicion int) int {
	return posicion*2 + 2
}
func calcularPosicionPadre(posicion int) int {
	return (posicion - 1) / 2
}
func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *colaConPrioridad[T]) Encolar(elemento T) {
	heap.datos = append(heap.datos, elemento)
	heap.cant++
	upheap(heap.datos, heap.cant, heap.cmp)
}

func upheap[T any](datos []T, cantidad int, funcion_cmp func(T, T) int) {
	posicionHijo := cantidad
	posicionPadre := calcularPosicionPadre(posicionHijo)
	padre := datos[posicionPadre]
	hijo := datos[posicionHijo]
	for funcion_cmp(hijo, padre) > 0 || posicionHijo != 0 {
		posicionHijo = posicionPadre
		posicionPadre = calcularPosicionPadre(posicionHijo)
	}
}

func (heap *colaConPrioridad[T]) VerMax() T {
	heap.verifcarColaVacia()
	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Desencolar() T {
	heap.verifcarColaVacia()
	dato := heap.datos[0]
	heap.cant--
	heap.datos[0] = heap.datos[heap.cant]
	heap.datos = heap.datos[:heap.cant]
	downheap(heap.datos, heap.cant, 0, heap.cmp)
	return dato
}

func downheap[T any](datos []T, cantidad int, posicion int, funcion_cmp func(T, T) int) {
	for posicion < cantidad {
		hijoIzquierdo := calcularPosicionHijoIzquierdo(posicion)
		hijoDerecho := calcularPosicionHijoDerecho(posicion)
		mayor := posicion

		if hijoIzquierdo < cantidad && funcion_cmp(datos[hijoIzquierdo], datos[mayor]) > 0 {
			mayor = hijoIzquierdo
		}

		if hijoDerecho < cantidad && funcion_cmp(datos[hijoDerecho], datos[mayor]) > 0 {
			mayor = hijoDerecho
		}

		if mayor == posicion {
			break
		}

		datos[posicion], datos[mayor] = datos[mayor], datos[posicion]
		posicion = mayor
	}
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func (heap *colaConPrioridad[T]) verifcarColaVacia() {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func Heapify[T any](elementos []T, cant int, funcion_cmp func(T, T) int) {
	for i := cant; i >= 0; i-- {
		downheap(elementos, len(elementos), i, funcion_cmp)
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	Heapify(elementos, len(elementos), funcion_cmp)
	for i := len(elementos) - 1; i >= 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downheap(elementos, i, 0, funcion_cmp)
	}
}
