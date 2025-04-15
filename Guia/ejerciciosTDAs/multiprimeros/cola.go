package multiprimeros

type Cola[T any] interface {

	// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
	EstaVacia(password int) bool

	// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje
	// "La cola esta vacia".
	VerPrimero(password int) T

	// Encolar agrega un nuevo elemento a la cola, al final de la misma.
	Encolar(dato T, password int)

	// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
	Desencolar(password int) T

	Multiprimeros(k int) []T
}
