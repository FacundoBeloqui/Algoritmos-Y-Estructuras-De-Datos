package ejerciciosTDAs

import "fmt"

/*
Implementar el TDA Fracción. Dicho TDA debe tener las siguientes primitivas, cuya documentación puede encontrarse en fraccion.go:

CrearFraccion(numerador, denominador int) Fraccion
Sumar(otra Fraccion) Fraccion
Multiplicar(otra Fraccion) Fraccion
ParteEntera() int
Representacion() string

Nota: considerar que se puede utilizar la función del módulo fmt Sprintf para generar la representación de la fracción.
*/
type Fraccion struct {
	num int
	den int
}

// CrearFraccion crea una fraccion con el numerador y denominador indicados.
// Si el denominador es 0, entra en panico.
func CrearFraccion(numerador, denominador int) Fraccion {
	if denominador == 0 {
		panic("el denominador no puede ser 0")
	} else {
		return Fraccion{
			num: numerador,
			den: denominador,
		}
	}
}

// Sumar crea una nueva fraccion, con el resultante de hacer la suma de las fracciones originales
func (f Fraccion) Sumar(otra Fraccion) Fraccion {
	nuevoNumerador := f.num*otra.den + otra.num*f.den
	nuevoDenominador := f.den * otra.den
	return Fraccion{num: nuevoNumerador, den: nuevoDenominador}
}

// Multiplicar crea una nueva fraccion con el resultante de multiplicar ambas fracciones originales
func (f Fraccion) Multiplicar(otra Fraccion) Fraccion {
	return Fraccion{num: f.num * otra.num, den: f.den * otra.den}
}

// ParteEntera devuelve la parte entera del numero representado por la fracción.
// Por ejemplo, para "7/2" = 3.5 debe devolver 3.
func (f Fraccion) ParteEntera() int {
	return f.num / f.den
}

// Representacion devuelve una representación en cadena de la fraccion simplificada (por ejemplo, no puede devolverse
// "10/8" sino que debe ser "5/4"). Considerar que si se trata de un número entero, debe mostrarse como tal.
// Considerar tambien el caso que se trate de un número negativo.
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func (f Fraccion) Representacion() string {
	divisor := gcd(f.num, f.den)
	f.num /= divisor
	f.den /= divisor
	if f.num < 0 && f.den < 0 {
		f.num *= -1
		f.den *= -1
	}
	return fmt.Sprintf("%d/%d", f.num, f.den)
}
