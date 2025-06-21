package algueiza

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func crearVuelo(datos []string) vuelo {
	numeroVuelo, _ := strconv.Atoi(datos[0])
	fecha := datos[6]
	prioridad, _ := strconv.Atoi(datos[5])
	atraso, _ := strconv.Atoi(datos[7])
	tiempoDeVuelo, _ := strconv.Atoi(datos[8])
	cancelado, _ := strconv.Atoi(datos[9])

	vueloProcesado := vuelo{
		numeroVuelo:   numeroVuelo,
		aerolinea:     datos[1],
		origen:        datos[2],
		destino:       datos[3],
		matricula:     datos[4],
		prioridad:     prioridad,
		fecha:         fecha,
		atraso:        atraso,
		tiempoDeVuelo: tiempoDeVuelo,
		cancelado:     cancelado,
	}
	return vueloProcesado
}
func leerArchivo(archivo string, tablero Tablero) error {
	file, err := os.Open(archivo)
	if err != nil {
		return fmt.Errorf("Error en comando agregar_archivo")
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		linea := s.Text()
		datos := strings.Split(linea, ",")
		vuelo := crearVuelo(datos)
		tablero.AgregarVuelo(vuelo)
	}
	return nil
}

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
		err := leerArchivo(comando[1], tablero)
		if err != nil {
			return fmt.Errorf("Error en comando agregar_archivo")
		}
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
