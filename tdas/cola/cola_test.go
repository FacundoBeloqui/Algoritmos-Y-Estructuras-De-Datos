package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "me fijo que no pueda ver el primero una cola vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "me fijo que no pueda desencolar una cola vacia")
	require.True(t, cola.EstaVacia())
}

func TestColaComprobarFIFO(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)
	cola.Encolar(5)
	require.EqualValues(t, cola.Desencolar(), 1)
	require.EqualValues(t, cola.Desencolar(), 2)
	require.EqualValues(t, cola.Desencolar(), 3)
	require.EqualValues(t, cola.Desencolar(), 4)
	require.EqualValues(t, cola.Desencolar(), 5)
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "me fijo que no pueda desencolar una cola vacia")
}
func TestPruebaDeVolumenCola(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i < 10000; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < 10000; i++ {
		require.EqualValues(t, cola.VerPrimero(), i)
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia())
}
func TestDesencolarHastaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "me fijo que no pueda ver el primero una cola vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "me fijo que no pueda desencolar una cola vacia")
	require.True(t, cola.EstaVacia())
}
func TestEncolarEnteros(t *testing.T) {
	colaEnteros := TDACola.CrearColaEnlazada[int]()
	colaEnteros.Encolar(1)
	require.EqualValues(t, colaEnteros.VerPrimero(), 1)
}
func TestEncolarStrings(t *testing.T) {
	colaCadenas := TDACola.CrearColaEnlazada[string]()
	colaCadenas.Encolar("Hola")
	require.EqualValues(t, colaCadenas.VerPrimero(), "Hola")
}
func TestEncolarFloats(t *testing.T) {
	colaFloats := TDACola.CrearColaEnlazada[float64]()
	colaFloats.Encolar(3.14)
	require.EqualValues(t, colaFloats.VerPrimero(), 3.14)
}
