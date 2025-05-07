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
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("A") })
}

func TestAbbClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, sigue sin existir")

	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, abb.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("") })

	abbNum := TDADiccionario.CrearABB[int, string](cmpInt)
	require.False(t, abbNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Borrar(0) })
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
	claves := []int{6, 12, 1}
	valores := []string{"valor1", "valor2", "valor3"}
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
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
	abb := TDADiccionario.CrearABB[int, string](cmpInt)
	abb.Guardar(8, "valor1")
	abb.Guardar(15, "valor2")
	require.True(t, abb.Pertenece(8))
	require.True(t, abb.Pertenece(15))
	require.EqualValues(t, "valor1", abb.Obtener(8))
	require.EqualValues(t, "valor2", abb.Obtener(15))
	require.EqualValues(t, 2, abb.Cantidad())

	abb.Guardar(8, "valor3")
	abb.Guardar(15, "valor4")
	require.True(t, abb.Pertenece(8))
	require.True(t, abb.Pertenece(15))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, "valor3", abb.Obtener(8))
	require.EqualValues(t, "valor4", abb.Obtener(15))
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

func TestAbbBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	claves := []int{15, 25, 10, 20, 30}
	valores := []string{"valor1", "valor2", "valor3", "valor4", "valor5"}
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	abb.Guardar(claves[3], valores[3])
	abb.Guardar(claves[4], valores[4])

	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], abb.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[1]) })
	require.EqualValues(t, 4, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[1]))

	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[0]) })
	require.EqualValues(t, 3, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[0]) })

	require.True(t, abb.Pertenece(claves[3]))
	require.EqualValues(t, valores[3], abb.Borrar(claves[3]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[3]) })
	require.EqualValues(t, 2, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[3]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[3]) })

	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[2]) })
	require.EqualValues(t, valores[4], abb.Borrar(claves[4]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[4]) })
	require.False(t, abb.Pertenece(claves[2]))
	require.False(t, abb.Pertenece(claves[4]))
	require.EqualValues(t, 0, abb.Cantidad())

}
