package calculadoraPolaca

import (
	"strings"
	"tdas/cola"
	"tdas/pila"
)

type Asociatividad struct {
	izquierda bool
}
type OperacionImpl struct {
	simbolo   string
	prioridad int
	asoc      Asociatividad
}

func (o OperacionImpl) Simbolo() string {
	return o.simbolo
}

func (o OperacionImpl) Prioridad() int {
	return o.prioridad
}

func (o OperacionImpl) Asociatividad() Asociatividad {
	return o.asoc
}

var Operaciones = []Operacion{
	OperacionImpl{simbolo: "+", prioridad: 2, asoc: Asociatividad{izquierda: true}},
	OperacionImpl{simbolo: "-", prioridad: 2, asoc: Asociatividad{izquierda: true}},
	OperacionImpl{simbolo: "/", prioridad: 3, asoc: Asociatividad{izquierda: true}},
	OperacionImpl{simbolo: "*", prioridad: 3, asoc: Asociatividad{izquierda: true}},
	OperacionImpl{simbolo: "^", prioridad: 4, asoc: Asociatividad{izquierda: false}},
	OperacionImpl{simbolo: "(", prioridad: 0, asoc: Asociatividad{izquierda: true}},
	OperacionImpl{simbolo: ")", prioridad: 0, asoc: Asociatividad{izquierda: true}},
}

func SepararCadena(cadena string) []string {
	operadores := []string{"(", ")", "+", "-", "*", "/", "^"}
	for _, operador := range operadores {
		cadena = strings.ReplaceAll(cadena, operador, " "+operador+" ")
	}
	return strings.Fields(cadena)
}

func EsOperador(o Operacion) bool {
	return o.Prioridad() != 0
}
func VerOperacion(cadena string) Operacion {
	for _, operador := range Operaciones {
		if cadena == operador.Simbolo() {
			return operador
		}
	}
	return OperacionImpl{cadena, 0, Asociatividad{true}}
}

func manejarOperador(cola cola.Cola[string], pila pila.Pila[Operacion], operador Operacion) {
	for !pila.EstaVacia() && (pila.VerTope().Prioridad() > operador.Prioridad() ||
		(pila.VerTope().Prioridad() == operador.Prioridad() && operador.Asociatividad().izquierda)) {
		cola.Encolar(pila.Desapilar().Simbolo())
	}
	pila.Apilar(operador)
}

func manejarParentesisApertura(pila pila.Pila[Operacion], parentesis Operacion) {
	pila.Apilar(parentesis)
}

func manejarParentesisCierre(cola cola.Cola[string], pila pila.Pila[Operacion]) {
	for !pila.EstaVacia() && pila.VerTope().Simbolo() != "(" {
		cola.Encolar(pila.Desapilar().Simbolo())
	}
	if !pila.EstaVacia() {
		pila.Desapilar()
	}
}

func manejarNumero(cola cola.Cola[string], token string) {
	cola.Encolar(token)
}

func ManejarToken(cola cola.Cola[string], pila pila.Pila[Operacion], token string) {
	if len(token) == 1 {
		caracter := VerOperacion(token)

		if EsOperador(caracter) {
			manejarOperador(cola, pila, caracter)
			return
		}

		if caracter.Simbolo() == "(" {
			manejarParentesisApertura(pila, caracter)
			return
		}

		if caracter.Simbolo() == ")" {
			manejarParentesisCierre(cola, pila)
			return
		}
	}
	manejarNumero(cola, token)
}

func VaciarPilaRestante(cola cola.Cola[string], pila pila.Pila[Operacion]) {
	for !pila.EstaVacia() {
		cola.Encolar(pila.Desapilar().Simbolo())
	}
}
