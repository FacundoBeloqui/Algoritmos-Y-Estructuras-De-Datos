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
		fecha:         datos[6],
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
	comandos := strings.Fields(linea)

	switch comandos[0] {
	case "agregar_archivo":
		return handleAgregarArchivo(comandos, tablero)

	case "ver_tablero":
		return handleVerTablero(comandos, tablero)

	case "info_vuelo":
		return handleInfoVuelo(comandos, tablero)

	case "prioridad_vuelos":
		return handlePrioridadVuelos(comandos, tablero)

	case "siguiente_vuelo":
		return handleSiguienteVuelo(comandos, tablero)

	case "borrar":
		return handleBorrar(comandos, tablero)

	default:
		return fmt.Errorf("Comando no valido")
	}
}

func handleAgregarArchivo(comandos []string, tablero Tablero) error{
	if len(comandos) != 2 {
		return fmt.Errorf("Error en comando agregar_archivo")
	}
	if _, err := os.Stat(comandos[1]); err != nil {
		return fmt.Errorf("Error en comando agregar_archivo")
	}
	
	if err := leerArchivo(comandos[1], tablero); err != nil {
		return fmt.Errorf("Error en comando agregar_archivo")
	}
	return nil
}

func handleVerTablero(comandos []string, tablero Tablero) error{
	if len(comandos) != 5 {
		return fmt.Errorf("Error en comando ver_tablero")
	}
	k, err := strconv.Atoi(comandos[1])
	if err != nil {
		return fmt.Errorf("Error en comando ver_tablero")
	}
	salidas, err := tablero.VerTablero(k, comandos[2], comandos[3], comandos[4])
	if err != nil {
		return err
	}
	for _, linea := range salidas {
		fmt.Println(strings.Join(linea, " - "))
	}
	return nil
}

func handleInfoVuelo(comandos []string, tablero Tablero) error{
	if len(comandos) != 2 {
		return fmt.Errorf("Error en comando info_vuelo")
	}
	codigo, err := strconv.Atoi(comandos[1])
	if err != nil {
		return fmt.Errorf("Error en comando info_vuelo")
	}
	info, err := tablero.InfoVuelo(codigo)
	if err != nil {
		return err
	}
	fmt.Println(strings.Join(info, " "))
	return nil
}

func handlePrioridadVuelos(comandos []string, tablero Tablero) error{
	if len(comandos) != 2 {
		return fmt.Errorf("Error en comando prioridad_vuelos")
	}
	k, err := strconv.Atoi(comandos[1])
	if err != nil || k < 0 {
		return fmt.Errorf("Error en comando prioridad_vuelos")
	}
	salidas := tablero.PrioridadVuelos(k)
	for _, linea := range salidas {
		fmt.Println(linea)
	}
	return nil
}

func handleSiguienteVuelo(comandos []string, tablero Tablero) error{
	if len(comandos) != 4 {
		return fmt.Errorf("Error en comando siguiente_vuelo")
	}
	info, encontrado := tablero.SiguienteVuelo(comandos[1], comandos[2], comandos[3])
	if !encontrado {
		fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s", comandos[1], comandos[2], comandos[3])
	}
	fmt.Println(strings.Join(info, " "))
	return nil
}

func handleBorrar(comandos []string, tablero Tablero) error{
	if len(comandos) != 3 {
		return fmt.Errorf("Error en comando borrar")
	}
	salidas, err := tablero.Borrar(comandos[1], comandos[2])
	if err != nil {
		return err
	}
	for _, linea := range salidas {
		fmt.Println(linea)
	}
	return nil
}