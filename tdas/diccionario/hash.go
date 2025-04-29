package diccionario

import(
	TDALista "tdas/lista"
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}
type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K,V]]
	tam      int
	cantidad int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {

}

func (h * hashAbierto[K, V]) Guardar(clave K, dato V){

}

func (h * hashAbierto[K, V]) Pertenece(clave K) bool{

}

func (h * hashAbierto[K, V]) Obtener(clave K) V{

}

func (h * hashAbierto[K, V]) Borrar(clave K) V{

}

func (h * hashAbierto[K, V]) Cantidad() int{

}

func (h * hashAbierto[K, V]) Iterar(func(clave K, dato V) bool){

}

func (h * hashAbierto[K, V]) Iterador() IterDiccionario[K, V]{

}

func (h * hashAbierto[K, V]) Iterador() IterDiccionario[K, V]{

}

//primitivas del iterador
// hay que hacer el struct del iterador del hash
func (i *) HaySiguiente() bool{

}

func (i *) VerActual() (K, V){

}

func (i *) Siguiente(){

}


