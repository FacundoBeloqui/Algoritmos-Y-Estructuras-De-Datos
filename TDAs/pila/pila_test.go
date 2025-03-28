package pila_test

import (
	TDAPila "TDAs/pila"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	// mas pruebas para este caso...
}
