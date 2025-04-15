package composicion

type Composicion struct {
	funciones []func(float64) float64
}

// CrearComposicion crea una nueva composición vacía
func CrearComposicion() Composicion {
	return Composicion{
		funciones: []func(float64) float64{},
	}
}

// AgregarFuncion agrega una función a la composición
func (c *Composicion) AgregarFuncion(f func(float64) float64) {
	c.funciones = append(c.funciones, f)
}

// Aplicar aplica las funciones en orden correcto (última agregada es la primera que se aplica)
func (c Composicion) Aplicar(x float64) float64 {
	for i := len(c.funciones) - 1; i >= 0; i-- {
		x = c.funciones[i](x)
	}
	return x
}
