package sonisomorficas

import (
	"tdas/diccionario"
)

/*
Dos cadenas X e Y son isomórficas si existe alguna transformación biyectiva de caracteres que permita obtener Y a partir
de X. Ejemplos: casa y bata son isomórficas, y la transformación es c → b, a → a, s → t. burro y pizza son isomórficas,
y la transformación es b → p, u → i, r → z, o → a. mesa y masa no son isomórficas, porque la transformación debe
ser biyectiva: no podemos incluir e → a y a → a. Escribir una función que reciba dos cadenas y determine si son
isomórficas. Indicar y justificar la complejidad de dicha función.
*/
func SonIsomorficas(cadena1, cadena2 string) bool {
	if len(cadena1) == len(cadena2) {
		dicc1 := diccionario.CrearHash[byte, byte]()
		dicc2 := diccionario.CrearHash[byte, byte]()
		for i := 0; i < len(cadena1); i++ {
			if !dicc1.Pertenece(cadena1[i]) {
				dicc1.Guardar(cadena1[i], cadena2[i])
			} else {
				if dicc1.Obtener(cadena1[i]) != cadena2[i] {
					return false
				}
			}
		}
		for i := 0; i < len(cadena2); i++ {
			if !dicc2.Pertenece(cadena2[i]) {
				dicc2.Guardar(cadena2[i], cadena1[i])
			} else {
				if dicc2.Obtener(cadena2[i]) != cadena1[i] {
					return false
				}
			}
		}
		return true
	}
	return false
}