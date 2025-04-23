package lista_test

import (
	TDALIsta "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIteradorInternoSuma(t *testing.T) {
	lista := TDALIsta.CrearListaEnlazada[int]()
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(8)
	lista.InsertarPrimero(10)

	suma := 0
	lista.Iterar(func(num int) bool {
		suma += num
		return true
	})

	require.Equal(t, 20, suma)
}

func TestIteradorInternoContar(t *testing.T) {
	lista := TDALIsta.CrearListaEnlazada[int]()
	lista.InsertarPrimero(22)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(11)
	lista.InsertarPrimero(86)
	lista.InsertarPrimero(1)

	contador := 0
	lista.Iterar(func(dato int) bool {
		contador++
		return true
	})

	require.Equal(t, 5, contador)
}

func TestIteradorInternoSinCorte(t *testing.T) {
	lista := TDALIsta.CrearListaEnlazada[int]()
	lista.InsertarPrimero(22)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(17)
	lista.InsertarPrimero(4)

	suma := 0
	lista.Iterar(func(num int) bool {
		if num%2 == 0 {
			suma += num
		}
		return true
	})

	require.Equal(t, 26, suma)
}

func TestIteradorInternoConCorte(t *testing.T) {
	lista := TDALIsta.CrearListaEnlazada[int]()
	lista.InsertarUltimo(8)
	lista.InsertarUltimo(2)
	lista.InsertarPrimero(6)
	lista.InsertarUltimo(9)
	lista.InsertarUltimo(5)

	vistos := []int{}
	contador := 0
	lista.Iterar(func(dato int) bool {
		vistos = append(vistos, dato)
		contador++
		return contador < 3
	})

	esperado := []int{6, 8, 2}
	require.Equal(t, esperado, vistos)
}

func TestIteradorInternoCorteFalse(t *testing.T) {
	lista := TDALIsta.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(20)
	lista.InsertarUltimo(30)
	lista.InsertarUltimo(40)

	vistos := []int{}
	lista.Iterar(func(num int) bool {
		vistos = append(vistos, num)
		return num != 30
	})

	esperado := []int{10, 20, 30}
	require.Equal(t, esperado, vistos)
}
func TestSeInsertaAlPrincipioAlCrearIteradorExterno(t *testing.T) {
	lista := TDALIsta.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	iterador.Insertar(1)
	require.Equal(t, lista.VerPrimero(), 1)
}

func TestInsertarAlFinalIteradorExterno(t *testing.T) {
	lista := TDALIsta.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}
	iterador.Insertar(3)
	require.Equal(t, 3, lista.VerUltimo())
}
func TestInsertarEnElMedioIteradorExterno(t *testing.T) {
	lista := TDALIsta.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	iterador := lista.Iterador()
	contador := 0
	largoOriginal := lista.Largo()

	for iterador.HaySiguiente() {
		if contador == largoOriginal/2 {
			iterador.Insertar(3)
		}
		contador++
		iterador.Siguiente()
	}

	iterador = lista.Iterador()
	var arr []int
	for iterador.HaySiguiente() {
		arr = append(arr, iterador.VerActual())
		iterador.Siguiente()
	}

	require.Equal(t, []int{1, 2, 3, 4, 5}, arr)
}

func TestBorrarAlCrearIteradorExterno(t *testing.T) {
	lista := TDALIsta.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	iterador := lista.Iterador()
	iterador.Borrar()
	require.Equal(t, 2, lista.VerPrimero())
}
func TestBorrarUltimoCambiaUltimoIteradorExterno(t *testing.T) {
	lista := TDALIsta.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	for iterador.HaySiguiente() {
		if iterador.VerActual() == lista.VerUltimo() {
			iterador.Borrar()
		}
		iterador.Siguiente()
	}
	require.Equal(t, 2, lista.VerUltimo())
}
func TestRemoverMedioIteradorExterno(t *testing.T) {
	lista := TDALIsta.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)

	iterador := lista.Iterador()
	contador := 0
	largoOriginal := lista.Largo()

	for iterador.HaySiguiente() {
		if contador == largoOriginal/2 {
			iterador.Borrar()
		}
		contador++
		iterador.Siguiente()
	}

	iterador = lista.Iterador()
	var arr []int
	for iterador.HaySiguiente() {
		arr = append(arr, iterador.VerActual())
		iterador.Siguiente()
	}

	require.Equal(t, []int{1, 2, 4, 5}, arr)
}
