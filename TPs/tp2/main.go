package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/diccionario"
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
	diccFechas := diccionario.CrearABB[string, algueiza.VueloImpl](strings.Compare)
	diccNumerosVuelo := diccionario.CrearHash[int, algueiza.VueloImpl]()

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
			algueiza.AgregarArchivo(comando[1], diccFechas, diccNumerosVuelo)
			println()

		case "ver_tablero":
			k, _ := strconv.Atoi(comando[1])
			algueiza.VerTablero(k, comando[2], comando[3], comando[4], diccFechas)
			println()

		case "info_vuelo":
			codigo, _ := strconv.Atoi(comando[1])
			algueiza.InfoVuelo(codigo, diccNumerosVuelo)
			println()

		case "prioridad_vuelo":
			k, _ := strconv.Atoi(comando[1])
			algueiza.PrioridadVuelos(k, diccNumerosVuelo)
			println()

		case "siguiente_vuelo":
			algueiza.SiguienteVuelo(comando[1], comando[2], comando[3], diccFechas, diccNumerosVuelo)
			println()

		case "borrar":
			algueiza.Borrar(comando[1], comando[2], diccFechas, diccNumerosVuelo)
			println()

		default:
			fmt.Println("Comando no valido, vuelva a ingresar una de las opciones mostradas")
			println()
		}
		mostrarComandos()
		println()
	}
	err := s.Err()
	if err != nil {
		fmt.Println(err)
	}

}
