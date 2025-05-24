package cola_prioridad_test

import (
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

var cmpInt = func(a, b int) int {
	return a - b
}

func TestColaPrioridadVacia(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestUnElemento(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	heap.Encolar(5)
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 5, heap.VerMax())
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 5, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestEsMaxHeap(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	arr := []int{9, 3, 5, 1, 6, 4, 5, 10, 21}

	for _, elem := range arr {
		heap.Encolar(elem)
	}
	anterior := heap.Desencolar()
	for !heap.EstaVacia() {
		actual := heap.Desencolar()
		require.LessOrEqual(t, actual, anterior)
		anterior = actual
	}
}

func TestEncolarYDesencolarStrings(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(strings.Compare)
	arr := []string{"A", "P", "C", "U", "Y"}

	heap.Encolar(arr[0])
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, "A", heap.VerMax())
	require.False(t, heap.EstaVacia())

	heap.Encolar(arr[1])
	require.EqualValues(t, 2, heap.Cantidad())
	require.EqualValues(t, "P", heap.VerMax())

	heap.Encolar(arr[2])
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, "P", heap.VerMax())
	require.EqualValues(t, "P", heap.Desencolar())
	require.EqualValues(t, "C", heap.VerMax())

	heap.Encolar(arr[3])
	require.EqualValues(t, 3, heap.Cantidad())
	require.EqualValues(t, "U", heap.VerMax())

	heap.Encolar(arr[4])
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, "Y", heap.VerMax())

	require.EqualValues(t, "Y", heap.Desencolar())
	require.EqualValues(t, "U", heap.VerMax())
	require.EqualValues(t, 3, heap.Cantidad())

	require.EqualValues(t, "U", heap.Desencolar())
	require.EqualValues(t, "C", heap.Desencolar())
	require.EqualValues(t, "A", heap.Desencolar())

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestEncolarYDesencolarEnteros(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	arr := []int{12, 24, 16, 30, 5}

	heap.Encolar(arr[0])
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 12, heap.VerMax())

	heap.Encolar(arr[1])
	require.EqualValues(t, 24, heap.VerMax())

	heap.Encolar(arr[2])
	heap.Encolar(arr[3])
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, 30, heap.VerMax())

	heap.Encolar(arr[4])
	require.EqualValues(t, 5, heap.Cantidad())

	require.EqualValues(t, 30, heap.Desencolar())
	require.EqualValues(t, 24, heap.VerMax())

	require.EqualValues(t, 24, heap.Desencolar())
	require.EqualValues(t, 16, heap.Desencolar())
	require.EqualValues(t, 12, heap.Desencolar())
	require.EqualValues(t, 5, heap.Desencolar())

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarYDesencolarFlotantes(t *testing.T) {
	cmpFloat := func(a, b float64) int {
		if a < b {
			return -1
		} else if a > b {
			return 1
		}
		return 0
	}

	heap := TDAColaPrioridad.CrearHeap(cmpFloat)
	arr := []float64{3.1, 2.5, 4.7, 1.9, 3.8}

	for _, val := range arr {
		heap.Encolar(val)
	}

	arrOrdenado := []float64{4.7, 3.8, 3.1, 2.5, 1.9}

	for _, val := range arrOrdenado {
		require.EqualValues(t, val, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestRepetidos(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	arr := []int{15, 6, 27, 27, 15}

	for _, val := range arr {
		heap.Encolar(val)
	}

	require.EqualValues(t, 27, heap.VerMax())
	require.EqualValues(t, 5, heap.Cantidad())

	require.EqualValues(t, 27, heap.Desencolar())
	require.EqualValues(t, 27, heap.Desencolar())
	require.EqualValues(t, 15, heap.VerMax())
	require.EqualValues(t, 15, heap.Desencolar())
	require.EqualValues(t, 15, heap.Desencolar())
	require.EqualValues(t, 6, heap.VerMax())
	require.EqualValues(t, 6, heap.Desencolar())

	require.True(t, heap.EstaVacia())
}

func TestHeapify(t *testing.T) {
	arr := []int{5, 1, 7, 2, 3, 6, 4, 8}
	heap := TDAColaPrioridad.CrearHeapArr(arr, cmpInt)

	require.EqualValues(t, 8, heap.VerMax())
	require.EqualValues(t, 8, heap.VerMax())

	for i := range arr {
		require.LessOrEqual(t, 8-i, heap.Desencolar())
	}
}

func TestCrearArregloVacio(t *testing.T) {
	var arr []int
	heap := TDAColaPrioridad.CrearHeapArr(arr, cmpInt)

	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapSort(t *testing.T) {
	arr := []int{5, 8, 1, 7, 20, 14, 24, 2}
	TDAColaPrioridad.HeapSort(arr, cmpInt)
	esperado := []int{1, 2, 5, 7, 8, 14, 20, 24}
	require.Equal(t, esperado, arr)
}

func TestVolumen(t *testing.T) {
	heap := TDAColaPrioridad.CrearHeap(cmpInt)
	const n = 1000000

	for i := range n {
		heap.Encolar(i)
		require.EqualValues(t, i, heap.VerMax())
	}

	anterior := heap.Desencolar()
	for !heap.EstaVacia() {
		actual := heap.Desencolar()
		require.LessOrEqual(t, actual, anterior)
		anterior = actual
	}
}
