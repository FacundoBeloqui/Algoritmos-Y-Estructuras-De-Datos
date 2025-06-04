package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tp2/algueiza"
)

func main() {
	tablero := algueiza.CrearTablero()
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		if s.Text() == "" {
			break
		}
		entrada := s.Text()
		comando := strings.Fields(entrada)
		switch comando[0] {
		case "agregar_archivo":
			if len(comando) != 2{
				fmt.Fprintln(os.Stderr, "Error en comando agregar_archivo")
				continue
			}
			//esto lo podriamos hacer de otra forma. basicamente verifico si se puede abrir el archivo, tipo si existe el archivo y bla. 
			if _, err := os.Stat(comando[1]); err != nil {
				fmt.Fprintln(os.Stderr, "Error en comando agregar_archivo")
				continue
			}
			tablero.AgregarArchivo(comando[1])

		case "ver_tablero":
			if len(comando) != 5{
				fmt.Fprintln(os.Stderr, "Error en comando ver_tablero")
				continue
			}
			k, _ := strconv.Atoi(comando[1])
			tablero.VerTablero(k, comando[2], comando[3], comando[4])

		case "info_vuelo":
			if len(comando) != 2{
				fmt.Fprintln(os.Stderr, "Error en comando info_vuelo")
				continue
			}
			codigo, _ := strconv.Atoi(comando[1])
			tablero.InfoVuelo(codigo)

		case "prioridad_vuelos":
			if len(comando) != 2{
				fmt.Fprintln(os.Stderr, "Error en comando prioridad_vuelos")
				continue
			}
			k, err := strconv.Atoi(comando[1])
			if err != nil || k < 0{
				fmt.Fprintln(os.Stderr, "Error en comando prioridad_vuelos")
				continue
			}
			tablero.PrioridadVuelos(k)

		case "siguiente_vuelo":
			if len(comando) != 4{
				fmt.Fprintln(os.Stderr, "Error en comando siguiente_vuelo")
				continue
			}
			tablero.SiguienteVuelo(comando[1], comando[2], comando[3])

		case "borrar":
			if len(comando) != 3{
				fmt.Fprintln(os.Stderr, "Error en comando borrar")
				continue
			}
			tablero.Borrar(comando[1], comando[2])

		default:
			fmt.Println("Comando no valido, vuelva a ingresar una de las opciones mostradas")
		}
		err := s.Err()
		if err == nil {
			fmt.Println("OK")
		} else {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
