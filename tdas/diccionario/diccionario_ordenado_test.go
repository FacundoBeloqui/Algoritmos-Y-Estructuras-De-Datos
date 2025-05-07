package diccionario_test

import (
	//"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var cmpInt = func(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func TestAbbVacio(t *testing.T) {
	t.Log("Comprueba que ABB vacío no tiene claves")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("A") })
	//require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("A") })
}

func TestAbbClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, sigue sin existir")

	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, abb.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("") })
	//require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("") })

	abbNum := TDADiccionario.CrearABB[int, string](cmpInt)
	require.False(t, abbNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Obtener(0) })
	//require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Borrar(0) })
}

func TestAbbUnElemento(t *testing.T) {
	t.Log("Comprueba que ABB con un elemento tiene esa Clave, unicamente")
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar("A", 10)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("A"))
	require.False(t, abb.Pertenece("B"))
	require.EqualValues(t, 10, abb.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("B") })
}

func TestAbbGuardar(t *testing.T) {
	t.Log("Guarda elementos en el ABB y comprueba que las claves y valores sean correctos en todo momento")
	claves := []string{"Gato", "Perro", "Vaca"}
	valores := []string{"miau", "guau", "moo"}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.False(t, abb.Pertenece(claves[1]))
	require.False(t, abb.Pertenece(claves[2]))

	abb.Guardar(claves[1], valores[1])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))
	require.False(t, abb.Pertenece(claves[2]))

	abb.Guardar(claves[2], valores[2])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, 3, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))
	require.EqualValues(t, valores[2], abb.Obtener(claves[2]))
}

func TestReemplazoDatos(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar("Gato", "miau")
	abb.Guardar("Perro", "guau")
	require.True(t, abb.Pertenece("Gato"))
	require.True(t, abb.Pertenece("Perro"))
	require.EqualValues(t, "miau", abb.Obtener("Gato"))
	require.EqualValues(t, "guau", abb.Obtener("Perro"))
	require.EqualValues(t, 2, abb.Cantidad())

	abb.Guardar("Gato", "miu")
	abb.Guardar("Perro", "baubau")
	require.True(t, abb.Pertenece("Gato"))
	require.True(t, abb.Pertenece("Perro"))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, "miu", abb.Obtener("Gato"))
	require.EqualValues(t, "baubau", abb.Obtener("Perro"))
}

func TestReemplazoDatosHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	abb := TDADiccionario.CrearABB[int, int](cmpInt)
	for i := 0; i < 500; i++ {
		abb.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		abb.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = abb.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}
