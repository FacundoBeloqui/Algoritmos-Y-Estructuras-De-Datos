package parte_entera_raiz

/*
Implementar un algoritmo que obtenga la parte entera de la raíz de un número n entero en O (log n). Justificar la
complejidad de la primitiva implementada.
*/
func Parte_entera(n int) int {
	return parte_entera_rec(n, 0, n)
}

func parte_entera_rec(n, inicio, fin int) int {
	if inicio > fin {
		return fin // cuando ya no hay más rango, fin es el último candidato válido
	}

	medio := (inicio + fin) / 2

	if medio*medio == n {
		return medio
	} else if medio*medio < n {
		return parte_entera_rec(n, medio+1, fin)
	} else {
		return parte_entera_rec(n, inicio, medio-1)
	}
}