package multiprimeros

type NodoCola[T any] struct {
	dato      T
	siguiente *NodoCola[T]
}

type ColaEnlazada[T any] struct {
	primero *NodoCola[T]
	ultimo  *NodoCola[T]
}
