package diccionario_test

import (
	"github.com/stretchr/testify/require"
	"math/rand/v2"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"
)

var cmpInt = func(a, b int) int {
	return a - b
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

func TestPeorCaso(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](cmpInt)
	n := 1000

	for i := range n {
		abb.Guardar(i, i)
		require.True(t, abb.Pertenece(i))
		require.Equal(t, i, abb.Obtener(i))
	}

	for i := range n {
		require.Equal(t, i, abb.Borrar(i))
		require.False(t, abb.Pertenece(i))
		require.Equal(t, n-i-1, abb.Cantidad())
	}

	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(0) })
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
func buscar(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteracionInorder(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	abb.Guardar(5, "E")
	abb.Guardar(2, "B")
	abb.Guardar(1, "A")
	abb.Guardar(6, "F")
	abb.Guardar(3, "C")
	abb.Guardar(4, "D")

	iter := abb.Iterador()
	claves := []int{1, 2, 3, 4, 5, 6}
	valores := []string{"A", "B", "C", "D", "E", "F"}

	for i := 0; iter.HaySiguiente(); i++ {
		clave, valor := iter.VerActual()
		require.Equal(t, claves[i], clave)
		require.Equal(t, valores[i], valor)
		iter.Siguiente()
	}
}

func TestIteradorInternoAbbClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	abb := TDADiccionario.CrearABB[string, *int](strings.Compare)
	abb.Guardar(claves[0], nil)
	abb.Guardar(claves[1], nil)
	abb.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	abb.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoAbbValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoAbbValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar(clave0, 7)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	abb.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoConRango(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](cmpInt)
	claves := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	for _, clave := range claves {
		abb.Guardar(clave, clave)
	}

	desde, hasta := 30, 70
	var esperado []int

	abb.IterarRango(&desde, &hasta, func(clave, valor int) bool {
		esperado = append(esperado, clave)
		return true
	})

	require.Equal(t, []int{30, 40, 50, 60, 70}, esperado)
}

func TestIteradorInternoConCondicionDeCorte(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](cmpInt)
	claves := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	for _, clave := range claves {
		abb.Guardar(clave, clave)
	}

	desde, hasta := 30, 70
	var esperado []int

	abb.IterarRango(&desde, &hasta, func(clave, valor int) bool {
		esperado = append(esperado, clave)
		return clave < 50
	})

	require.Equal(t, []int{30, 40, 50}, esperado)
}

func TestIteradorRango1(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	abb.Guardar(10, "A")
	abb.Guardar(5, "B")
	abb.Guardar(15, "C")
	abb.Guardar(3, "D")
	abb.Guardar(7, "E")
	abb.Guardar(12, "F")
	abb.Guardar(18, "G")

	desde := 6
	hasta := 15
	iter := abb.IteradorRango(&desde, &hasta)

	var claves []int
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves = append(claves, clave)
		iter.Siguiente()
	}

	esperado := []int{7, 10, 12, 15}
	require.Equal(t, esperado, claves)
}

func TestIteradorRango2(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	abb.Guardar(10, "A")
	abb.Guardar(5, "B")
	abb.Guardar(15, "C")
	abb.Guardar(3, "D")
	abb.Guardar(7, "E")
	abb.Guardar(12, "F")
	abb.Guardar(18, "G")

	iter := abb.IteradorRango(nil, nil)

	var claves []int
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves = append(claves, clave)
		iter.Siguiente()
	}

	esperado := []int{3, 5, 7, 10, 12, 15, 18}
	require.Equal(t, esperado, claves)
}

func TestIteradorRango3(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	abb.Guardar(10, "A")
	abb.Guardar(5, "B")
	abb.Guardar(15, "C")
	abb.Guardar(3, "D")
	abb.Guardar(7, "E")
	abb.Guardar(12, "F")
	abb.Guardar(18, "G")

	desde := 12
	iter := abb.IteradorRango(&desde, nil)

	var claves []int
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves = append(claves, clave)
		iter.Siguiente()
	}

	esperado := []int{12, 15, 18}
	require.Equal(t, esperado, claves)
}

func TestIteradorRango4(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	abb.Guardar(10, "A")
	abb.Guardar(5, "B")
	abb.Guardar(15, "C")
	abb.Guardar(3, "D")
	abb.Guardar(7, "E")
	abb.Guardar(12, "F")
	abb.Guardar(18, "G")

	hasta := 12
	iter := abb.IteradorRango(nil, &hasta)

	var claves []int
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves = append(claves, clave)
		iter.Siguiente()
	}

	esperado := []int{3, 5, 7, 10, 12}
	require.Equal(t, esperado, claves)
}

func TestAbbBorrarRaiz(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	abb.Guardar(11, "A")
	abb.Guardar(8, "B")
	abb.Guardar(15, "C")
	abb.Guardar(3, "D")
	abb.Guardar(5, "E")
	abb.Guardar(12, "F")
	abb.Guardar(18, "G")

	abb.Borrar(11)
	iter := abb.Iterador()

	var claves []int
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		claves = append(claves, clave)
		iter.Siguiente()
	}

	esperado := []int{3, 5, 8, 12, 15, 18}
	require.Equal(t, esperado, claves)
}

func TestAbbIteradorExterno(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, string](cmpInt)

	claves := []int{25, 10, 32, 14, 2, 1}
	valores := []string{"A", "B", "C", "D", "E", "F"}
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	abb.Guardar(claves[3], valores[3])
	abb.Guardar(claves[4], valores[4])
	abb.Guardar(claves[5], valores[5])

	clavesEsperadas := []int{1, 2, 10, 14, 25, 32}
	valoresEsperados := []string{"F", "E", "B", "D", "A", "C"}

	iter := abb.Iterador()

	for i := 0; iter.HaySiguiente(); i++ {
		clave, valor := iter.VerActual()
		require.Equal(t, clavesEsperadas[i], clave)
		require.Equal(t, valoresEsperados[i], valor)
		require.True(t, iter.HaySiguiente())
		iter.Siguiente()
	}

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })

}
func randRange(min, max int) int {
	return rand.IntN(max+1-min) + min
}

func TestAbbVolumen(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](cmpInt)
	const maximo = 1000000
	const minimo = 0

	for abb.Cantidad() != maximo {
		clave := randRange(minimo, maximo)
		if !abb.Pertenece(clave) {
			abb.Guardar(clave, clave)
		}
	}
	for iter := abb.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		abb.Borrar(clave)
	}
}
