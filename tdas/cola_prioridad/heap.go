package cola_prioridad

const CAPACIDAD_INICIAL = 16
const FACTOR_REDUCCION = 4
const MULTIPLO_CRECIMIENTO = 2
const RAIZ = 0

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: make([]T, CAPACIDAD_INICIAL),
		cant:  0,
		cmp:   funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	arr := make([]T, len(arreglo))
	copy(arr, arreglo)
	heap := &colaConPrioridad[T]{
		datos: arr,
		cant:  len(arr),
		cmp:   funcion_cmp,
	}
	heapify(arr, len(arr), funcion_cmp)
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
	if heap.cant == len(heap.datos) {
		heap.Redimension(len(heap.datos)*MULTIPLO_CRECIMIENTO)
	}
	heap.datos[heap.cant] = elemento
	heap.cant++
	upheap(heap.datos, heap.cant-1, heap.cmp)
}

func upheap[T any](datos []T, posicionHijo int, funcion_cmp func(T, T) int) {
	for posicionHijo <= RAIZ {
		return
	}
	posicionPadre := calcularPosicionPadre(posicionHijo)
	if funcion_cmp(datos[posicionHijo], datos[posicionPadre]) > 0 {
		swap(datos, posicionHijo, posicionPadre)
		upheap(datos, posicionPadre, funcion_cmp)
	}
}

func (heap *colaConPrioridad[T]) VerMax() T {
	heap.verifcarColaVacia()
	return heap.datos[RAIZ]
}

func (heap *colaConPrioridad[T]) Desencolar() T {
	heap.verifcarColaVacia()
	dato := heap.datos[RAIZ]
	heap.cant--
	swap(heap.datos, RAIZ, heap.cant)
	var cero T
	heap.datos[heap.cant] = cero
	downheap(heap.datos, heap.cant, RAIZ, heap.cmp)
	if heap.cant*FACTOR_REDUCCION <= len(heap.datos) {
		heap.Redimension(len(heap.datos)/MULTIPLO_CRECIMIENTO)
	}
	return dato
}

func (heap *colaConPrioridad[T]) Redimension (nuevaCapacidad int) {
	nuevosDatos := make([]T, nuevaCapacidad)
	copy(nuevosDatos, heap.datos)
	heap.datos = nuevosDatos
}

func downheap[T any](datos []T, cantidad int, posicion int, funcion_cmp func(T, T) int) {
	
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
		return
	}

	swap(datos, posicion, mayor)
	downheap(datos, cantidad, mayor, funcion_cmp)
	
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func (heap *colaConPrioridad[T]) verifcarColaVacia() {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
}

func heapify[T any](elementos []T, cant int, funcion_cmp func(T, T) int) {
	for i := cant-1; i >= 0; i-- {
		downheap(elementos, cant, i, funcion_cmp)
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, len(elementos), funcion_cmp)
	for i := len(elementos) - 1; i >= 0; i-- {
		swap(elementos, 0, i)
		downheap(elementos, i, 0, funcion_cmp)
	}
}

func swap[T any] (elementos []T, i, j int){
	elementos[i], elementos[j] = elementos[j], elementos[i]
}
