package algueiza

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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
		fmt.Printf("Error %v al abrir el archivo %s", archivo, err)
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
}

func VerTablero(k int, modo string, desde string, hasta string, dicc diccionario.DiccionarioOrdenado[string, VueloImpl]) {
	if k <= 0 || modo != "asc" && modo != "desc" || hasta < desde {
		panic("Vuelva a intentarlo")
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
}
func InfoVuelo(codigo int, diccionario2 diccionario.Diccionario[int, VueloImpl]) {
	if !diccionario2.Pertenece(codigo) {
		panic("El codigo no existe")
	}
	datos := diccionario2.Obtener(codigo)
	fmt.Printf("%d %s %s %s %s %d %s %d %d %d\n",
		datos.numeroVuelo, datos.aerolinea, datos.origen, datos.destino,
		datos.matricula, datos.prioridad, datos.fecha, datos.atraso, datos.tiempoDeVuelo, datos.cancelado)

}
