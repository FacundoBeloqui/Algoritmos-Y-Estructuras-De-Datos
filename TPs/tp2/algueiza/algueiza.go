package algueiza

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"tdas/pila"
)

type VueloImpl struct {
	numeroVuelo   int
	aerolinea     string
	origen        string
	destino       string
	matricula     string
	prioridad     int
	fecha         string
	atraso        int
	tiempoDeVuelo int
	cancelado     int
}

func AgregarArchivo(archivo string, diccFechas diccionario.DiccionarioOrdenado[string, VueloImpl], diccNumerosVuelo diccionario.Diccionario[int, VueloImpl]) {
	file, err := os.Open(archivo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error en comando agregar_archivo: no se pudo abrir el archivo %s: %v\n", archivo, err)
		return
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, linea := range data {
		numeroVuelo, _ := strconv.Atoi(linea[0])
		fecha := linea[6]
		prioridad, _ := strconv.Atoi(linea[5])
		atraso, _ := strconv.Atoi(linea[7])
		tiempoDeVuelo, _ := strconv.Atoi(linea[8])
		cancelado, _ := strconv.Atoi(linea[9])
		datos := VueloImpl{numeroVuelo: numeroVuelo, aerolinea: linea[1], origen: linea[2], destino: linea[3], matricula: linea[4], prioridad: prioridad, fecha: fecha, atraso: atraso, tiempoDeVuelo: tiempoDeVuelo, cancelado: cancelado}
		if diccNumerosVuelo.Pertenece(numeroVuelo) {
			datosAnterior := diccNumerosVuelo.Obtener(numeroVuelo)
			diccFechas.Borrar(datosAnterior.fecha)
			diccNumerosVuelo.Borrar(numeroVuelo)
		}
		diccNumerosVuelo.Guardar(numeroVuelo, datos)
		diccFechas.Guardar(fecha, datos)
	}
	println("OK")
}

func VerTablero(k int, modo string, desde string, hasta string, dicc diccionario.DiccionarioOrdenado[string, VueloImpl]) {
	if k <= 0 {
		fmt.Fprintf(os.Stderr, "Error en comando ver_tablero: cantidad no válida\n")
		return
	}

	if modo != "asc" && modo != "desc" {
		fmt.Fprintf(os.Stderr, "Error en comando ver_tablero: modo no válido\n")
		return
	}

	if hasta < desde {
		fmt.Fprintf(os.Stderr, "Error en comando ver_tablero: hasta es mayor que desde\n")
		return
	}
	pilaAux := pila.CrearPilaDinamica[VueloImpl]()
	contador := 0
	for iter := dicc.IteradorRango(&desde, &hasta); iter.HaySiguiente() && contador < k; iter.Siguiente() {
		_, valor := iter.VerActual()
		if modo == "desc" {
			pilaAux.Apilar(valor)
		} else {
			fmt.Printf("%s - %d\n", valor.fecha, valor.numeroVuelo)
		}
		contador++
	}
	for !pilaAux.EstaVacia() {
		valor := pilaAux.Desapilar()
		fmt.Printf("%s - %d\n", valor.fecha, valor.numeroVuelo)
	}
	println("OK")
}
func InfoVuelo(codigo int, diccionario2 diccionario.Diccionario[int, VueloImpl]) {
	if !diccionario2.Pertenece(codigo) {
		fmt.Printf("Error en comando info_vuelo: el vuelo %d no fue encontrado\n", codigo)
		return
	}
	datos := diccionario2.Obtener(codigo)
	fmt.Printf("%d %s %s %s %s %d %s %d %d %d\n",
		datos.numeroVuelo, datos.aerolinea, datos.origen, datos.destino,
		datos.matricula, datos.prioridad, datos.fecha, datos.atraso, datos.tiempoDeVuelo, datos.cancelado)
	println("OK")
}
func cmp(a, b VueloImpl) int {
	if a.prioridad > b.prioridad {
		return 1
	} else if a.prioridad < b.prioridad {
		return -1
	} else {
		if a.numeroVuelo > b.numeroVuelo {
			return 1
		} else if a.numeroVuelo < b.numeroVuelo {
			return -1
		}
	}
	return 0
}

func TopK(arr []VueloImpl, k int) []VueloImpl {
	cp := cola_prioridad.CrearHeapArr(arr, cmp)
	top := make([]VueloImpl, k)

	for i := 0; i < k; i++ {
		top[i] = cp.Desencolar()
	}
	return top
}
func PrioridadVuelos(k int, diccionario2 diccionario.Diccionario[int, VueloImpl]) {
	miarr := make([]VueloImpl, diccionario2.Cantidad())
	i := 0
	for iter := diccionario2.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, datos := iter.VerActual()
		miarr[i] = datos
		i++
	}
	topVuelos := TopK(miarr, k)
	for _, elem := range topVuelos {
		fmt.Printf("%d - %d\n", elem.prioridad, elem.numeroVuelo)
	}
	println("OK")
}

func SiguienteVuelo(origen, destino, fecha string, dicc diccionario.DiccionarioOrdenado[string, VueloImpl], diccionario2 diccionario.Diccionario[int, VueloImpl]) {
	encontrado := false
	for iter := dicc.IteradorRango(&fecha, nil); iter.HaySiguiente(); iter.Siguiente() {
		_, valor := iter.VerActual()
		if valor.destino == destino && valor.origen == origen {
			encontrado = true
			InfoVuelo(valor.numeroVuelo, diccionario2)
		}
	}
	if !encontrado {
		fmt.Printf("Error en comando siguiente_vuelo: no hay vuelo registrado desde %s hacia %s desde %s", origen, destino, fecha)
		println("OK")
	}
}

func Borrar(desde, hasta string, dicc diccionario.DiccionarioOrdenado[string, VueloImpl], diccionario2 diccionario.Diccionario[int, VueloImpl]) {
	if desde > hasta {
		fmt.Print("Error en comando borrar: hasta es mayor que desde")
		return
	}
	for iter := dicc.IteradorRango(&desde, &hasta); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		InfoVuelo(valor.numeroVuelo, diccionario2)
		dicc.Borrar(clave)
		diccionario2.Borrar(valor.numeroVuelo)
	}
	println("OK")
}
