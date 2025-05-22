package cola_prioridad_test
import (
	//"fmt"
	//"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	"testing"
	"github.com/stretchr/testify/require"
)

var cmpInt = func(a, b int) int {
	return a - b
}

func TestColaPrioridadVacia(t *testing.T){
	heap := TDAColaPrioridad.CrearHeap[int](cmpInt)
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
	require.EqualValues(t,0, heap.Cantidad())
}

func TestEsMaxHeap(t *testing.T){
	heap := TDAColaPrioridad.CrearHeap[int](cmpInt)
	arr := []int{9, 3, 5, 1, 6, 4, 5, 10, 21}

	heap.Encolar(arr[0])
	heap.Encolar(arr[1])
	heap.Encolar(arr[2])
	heap.Encolar(arr[3])
	//heap.Encolar(arr[4])
	//heap.Encolar(arr[5])
	//heap.Encolar(arr[6])
	//heap.Encolar(arr[7])
	//heap.Encolar(arr[8])


	/*for _, elem := range arr {
		heap.Encolar(elem)
	}
	//verifico que siempre salga un numero menor al actual y de esta forma me aseguro que el maximo siempre esta arriba
	/*anterior := heap.Desencolar()
	for !heap.EstaVacia(){
		actual := heap.Desencolar()
		if cmpInt(anterior, actual) < 0{
			t.Fatal("No cumple con el orden de un heap de maximos")
		}
		anterior = actual
	}*/
}

func TestHeapSort(t *testing.T){
	arr := []int{5, 8, 1, 7, 20 ,14, 24, 2}
	TDAColaPrioridad.HeapSort(arr, cmpInt)
	esperado := []int{1, 2, 5, 7, 8, 14, 20, 24}
	require.Equal(t, esperado, arr)
}