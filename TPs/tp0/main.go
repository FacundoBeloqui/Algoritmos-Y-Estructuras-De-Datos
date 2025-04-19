package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const RUTA1 = "archivo1.in"
const RUTA2 = "archivo2.in"

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
	fmt.Println("Ingrese un input")
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	fmt.Printf("Leí: %s\n", s.Text())
}
