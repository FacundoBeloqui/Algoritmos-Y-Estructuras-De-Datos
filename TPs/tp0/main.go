package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const ruta1 = "archivo1.in"
const ruta2 = "archivo2.in"

func leerArchivo(ruta string) []int {
	var vectorOrdenado []int
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s", ruta, err)
	}
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		numero, e := strconv.Atoi(s.Text())
		if e == nil {
			vectorOrdenado = append(vectorOrdenado, numero)
		}
	}
	err = s.Err()
	if err != nil {
		fmt.Println(err)
	}
	return vectorOrdenado
}

func main() {
	vector1 := leerArchivo(ruta1)
	vector2 := leerArchivo(ruta2)
	comparacion := ejercicios.Comparar(vector1, vector2)
	if comparacion == 1 {
		ejercicios.Seleccion(vector1)
		for i := 0; i < len(vector1); i++ {
			fmt.Printf("%d\n", vector1[i])
		}

	} else if comparacion == -1 {
		ejercicios.Seleccion(vector2)
		for i := 0; i < len(vector2); i++ {
			fmt.Printf("%d\n", vector2[i])
		}
	}
}
