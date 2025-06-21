package algueiza

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ProcesarComando(linea string, tablero Tablero) error {
	comando := strings.Fields(linea)

	if len(comando) == 0 {
		return fmt.Errorf("Comando vacío")
	}

	switch comando[0] {
	case "agregar_archivo":
		if len(comando) != 2 {
			return fmt.Errorf("Error en comando agregar_archivo")
		}
		if _, err := os.Stat(comando[1]); err != nil {
			return fmt.Errorf("Error en comando agregar_archivo")
		}
		tablero.AgregarArchivo(comando[1])
		return nil

	case "ver_tablero":
		if len(comando) != 5 {
			return fmt.Errorf("Error en comando ver_tablero")
		}
		k, err := strconv.Atoi(comando[1])
		if err != nil {
			return fmt.Errorf("Error en comando ver_tablero")
		}
		salidas, err := tablero.VerTablero(k, comando[2], comando[3], comando[4])
		if err != nil {
			return err
		}
		for _, linea := range salidas {
			fmt.Println(strings.Join(linea, " - "))
		}
		return nil

	case "info_vuelo":
		if len(comando) != 2 {
			return fmt.Errorf("Error en comando info_vuelo")
		}
		codigo, err := strconv.Atoi(comando[1])
		if err != nil {
			return fmt.Errorf("Error en comando info_vuelo")
		}
		info, err := tablero.InfoVuelo(codigo)
		if err != nil {
			return err
		}
		fmt.Println(strings.Join(info, " "))
		return nil

	case "prioridad_vuelos":
		if len(comando) != 2 {
			return fmt.Errorf("Error en comando prioridad_vuelos")
		}
		k, err := strconv.Atoi(comando[1])
		if err != nil || k < 0 {
			return fmt.Errorf("Error en comando prioridad_vuelos")
		}
		salidas := tablero.PrioridadVuelos(k)
		for _, linea := range salidas {
			fmt.Println(linea)
		}
		return nil

	case "siguiente_vuelo":
		if len(comando) != 4 {
			return fmt.Errorf("Error en comando siguiente_vuelo")
		}
		info, encontrado := tablero.SiguienteVuelo(comando[1], comando[2], comando[3])
		if !encontrado {
			fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s", comando[1], comando[2], comando[3])
		}
		fmt.Println(strings.Join(info, " "))
		return nil

	case "borrar":
		if len(comando) != 3 {
			return fmt.Errorf("Error en comando borrar")
		}
		salidas, err := tablero.Borrar(comando[1], comando[2])
		if err != nil {
			return err
		}
		for _, linea := range salidas {
			fmt.Println(linea)
		}

		return nil

	default:
		return fmt.Errorf("Comando no valido")
	}
}
