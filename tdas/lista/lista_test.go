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
