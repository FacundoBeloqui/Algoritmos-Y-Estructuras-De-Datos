package composicion

/*
(★★) ♠♠ Implementar en Go el TDA ComposiciónFunciones que emula la composición de funciones (i.e. f(g(h(x))).
Se debe definir la estructura del TDA, y las siguientes primitivas:

CrearComposicion() ComposicionFunciones
AgregarFuncion(func (float64) float64)
Aplicar(float64) float64

Considerar que primero se irán agregando las funciones como se leen, pero tener en cuenta el correcto orden de aplicación.
Por ejemplo: para emular f(g(x)), se debe hacer:

composicion.AgregarFuncion(f)
composicion.AgregarFuncion(g)
composicion.Aplicar(x)

Indicar el orden de las primitivas.
*/

type ComposicionFunciones interface {
	AgregarFuncion(func(float64) float64)
	Aplicar(float64) float64
}
