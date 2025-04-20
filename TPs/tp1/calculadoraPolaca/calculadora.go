package calculadoraPolaca

import (
	"strings"
	"tdas/cola"
	"tdas/pila"
)

const PRECEDENCIA_SUMA_RESTA = 2
const PRECEDENCIA_MULT_DIV = 3
const PRECEDENCIA_EXPONENCIAL = 4
const PRECEDENCIA_NUMERO = 0

type Caracter struct {
	precedencia int
	Simbolo     rune
}

func SepararCadena(cadena string) []string {
	operadores := []string{"(", ")", "+", "-", "*", "/", "^"}
	for _, operador := range operadores {
		cadena = strings.ReplaceAll(cadena, operador, " "+operador+" ")
	}
	return strings.Fields(cadena)
}

func DeterminarCaracter(c rune) Caracter {
	switch c {
	case '+':
		return Caracter{PRECEDENCIA_SUMA_RESTA, c}
	case '-':
		return Caracter{PRECEDENCIA_SUMA_RESTA, c}
	case '*':
		return Caracter{PRECEDENCIA_MULT_DIV, c}
	case '/':
		return Caracter{PRECEDENCIA_MULT_DIV, c}
	case '^':
		return Caracter{PRECEDENCIA_EXPONENCIAL, c}
	default:
		return Caracter{PRECEDENCIA_NUMERO, c}
	}
}

func EsOperador(c Caracter) bool {
	return c.precedencia != 0
}

func ManejarToken(cola cola.Cola[string], pila pila.Pila[Caracter], token string) {
	if len(token) == 1 {
		c := rune(token[0])
		car := DeterminarCaracter(c)

		if EsOperador(car) {
			for !pila.EstaVacia() && (pila.VerTope().precedencia > car.precedencia ||
				(pila.VerTope().precedencia == car.precedencia && car.Simbolo != '^')) {
				cola.Encolar(string(pila.Desapilar().Simbolo))
			}
			pila.Apilar(car)
			return
		}

		if c == '(' {
			pila.Apilar(Caracter{0, '('})
			return
		}
		if c == ')' {
			for !pila.EstaVacia() && pila.VerTope().Simbolo != '(' {
				cola.Encolar(string(pila.Desapilar().Simbolo))
			}
			if !pila.EstaVacia() {
				pila.Desapilar()
			}
			return
		}
	}
	cola.Encolar(token)
}

func VaciarPilaRestante(cola cola.Cola[string], pila pila.Pila[Caracter]) {
	for !pila.EstaVacia() {
		cola.Encolar(string(pila.Desapilar().Simbolo))
	}
}
