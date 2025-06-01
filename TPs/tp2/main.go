package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tp2/algueiza"
)

func mostrarComandos() {
	fmt.Println("Elija una de las siguientes opciones:")
	fmt.Println("- agregar_archivo <nombre_archivo> ")
	fmt.Println("- ver_tablero <K cantidad vuelos> <modo: asc/desc> <desde> <hasta>")
	fmt.Println("- info_vuelo <código vuelo>")
	fmt.Println("- prioridad_vuelos <K cantidad vuelos>")
	fmt.Println("- siguiente_vuelo <aeropuerto origen> <aeropuerto destino> <fecha>")
	fmt.Println("- borrar <desde> <hasta>")
	println()
}
func main() {
	tablero := algueiza.CrearTablero()
	fmt.Println("Bienvenido a Algueiza!")
	mostrarComandos()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		entrada := s.Text()
		comando := strings.Fields(entrada)
		switch comando[0] {
		case "agregar_archivo":
			tablero.AgregarArchivo(comando[1])

		case "ver_tablero":
			k, _ := strconv.Atoi(comando[1])
			tablero.VerTablero(k, comando[2], comando[3], comando[4])

		case "info_vuelo":
			codigo, _ := strconv.Atoi(comando[1])
			tablero.InfoVuelo(codigo)

		case "prioridad_vuelos":
			k, _ := strconv.Atoi(comando[1])
			tablero.PrioridadVuelos(k)

		case "siguiente_vuelo":
			tablero.SiguienteVuelo(comando[1], comando[2], comando[3])

		case "borrar":
			tablero.Borrar(comando[1], comando[2])

		default:
			fmt.Println("Comando no valido, vuelva a ingresar una de las opciones mostradas")
		}
		println()
		err := s.Err()
		if err == nil {
			fmt.Println("OK")
		} else {
			fmt.Println(err)
		}
		mostrarComandos()
	}

}
