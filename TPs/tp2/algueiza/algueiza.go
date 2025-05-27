package algueiza

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"tdas/diccionario"
)

type vuelo struct {
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

var cmpInt = func(a, b int) int {
	return a - b
}

func AgregarArchivo(archivo string) diccionario.Diccionario[int, []string] {
	dicc := diccionario.CrearHash[int, []string]()
	file, err := os.Open(archivo)
	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s", archivo, err)
		return dicc
	}
	defer file.Close()
	reader := csv.NewReader(file)
	data, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	for _, vuelo := range data {
		numeroVuelo, _ := strconv.Atoi(vuelo[0])
		dicc.Guardar(numeroVuelo, vuelo[1:])
	}
	return dicc
}
