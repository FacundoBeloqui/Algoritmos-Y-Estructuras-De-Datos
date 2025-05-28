package algueiza

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"tdas/diccionario"
	//"tdas/pila"
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

func AgregarArchivo(archivo string, diccFechas diccionario.DiccionarioOrdenado[string, VueloImpl], diccNumerosVuelo diccionario.Diccionario[int, string]) {
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
		if !diccNumerosVuelo.Pertenece(numeroVuelo) {
			diccNumerosVuelo.Guardar(numeroVuelo, fecha)
		} else {
			nuevaFecha := diccNumerosVuelo.Obtener(numeroVuelo)
			diccFechas.Borrar(nuevaFecha)
			diccNumerosVuelo.Guardar(numeroVuelo, fecha)

		}
		prioridad, _ := strconv.Atoi(linea[5])
		atraso, _ := strconv.Atoi(linea[7])
		tiempoDeVuelo, _ := strconv.Atoi(linea[8])
		cancelado, _ := strconv.Atoi(linea[9])
		datos := VueloImpl{numeroVuelo: numeroVuelo, aerolinea: linea[1], origen: linea[2], destino: linea[3], matricula: linea[4], prioridad: prioridad, fecha: fecha, atraso: atraso, tiempoDeVuelo: tiempoDeVuelo, cancelado: cancelado}
		diccFechas.Guardar(linea[6], datos)
	}
}
func VerTablero(k int, modo string, desde string, hasta string, dicc diccionario.DiccionarioOrdenado[string, VueloImpl]) {
	if k <= 0 || modo != "asc" || modo != "desc" || hasta < desde {
		panic("Vuelva a intentarlo")
	}
	//for iter := dicc.IteradorRango(&desde, &hasta)
}
