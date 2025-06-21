package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "me fijo que no pueda desapilar una pila vacia")
	require.True(t, pila.EstaVacia())
	pila.Apilar(1)
	require.False(t, pila.EstaVacia())
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "me fijo que no pueda desapilar una pila vacia")
}

func TestPilaVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "me fijo que no pueda ver el tope de una pila vacia")
	pila.Apilar(1)
	require.EqualValues(t, pila.VerTope(), 1)
	pila.Apilar(15)
	require.EqualValues(t, pila.VerTope(), 15)
	pila.Desapilar()
	require.EqualValues(t, pila.VerTope(), 1)
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "me fijo que no pueda ver el tope de una pila vacia")
}

func TestPilaComprobarLIFO(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)
	pila.Apilar(5)
	require.EqualValues(t, pila.Desapilar(), 5)
	require.EqualValues(t, pila.Desapilar(), 4)
	require.EqualValues(t, pila.Desapilar(), 3)
	require.EqualValues(t, pila.Desapilar(), 2)
	require.EqualValues(t, pila.Desapilar(), 1)
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "me fijo que no pueda desapilar una pila vacia")
}

func TestPruebaDeVolumenPila(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 10000; i++ {
		pila.Apilar(i)
		require.EqualValues(t, pila.VerTope(), i)
	}
	for i := 9999; i > -1; i-- {
		require.EqualValues(t, pila.VerTope(), i)
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}

func TestApilarDistintosTiposDeDato(t *testing.T) {
	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	pilaCadenas := TDAPila.CrearPilaDinamica[string]()
	pilaFloats := TDAPila.CrearPilaDinamica[float64]()
	pilaEnteros.Apilar(1)
	pilaCadenas.Apilar("Hola")
	pilaFloats.Apilar(3.14)
	require.EqualValues(t, pilaEnteros.VerTope(), 1)
	require.EqualValues(t, pilaCadenas.VerTope(), "Hola")
	require.EqualValues(t, pilaFloats.VerTope(), 3.14)
}
