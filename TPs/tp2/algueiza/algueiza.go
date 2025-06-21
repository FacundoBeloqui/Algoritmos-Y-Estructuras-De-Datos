package algueiza

import (
	"bufio"
	"errors"
	"tdas/cola_prioridad"
	"tdas/pila"

	"fmt"
	"os"
	"strconv"
	"strings"
	"tdas/diccionario"
)

type TableroImpl struct {
	vuelosCodigo diccionario.Diccionario[int, vuelo]
	vuelosFecha  diccionario.DiccionarioOrdenado[claveVuelo, vuelo]
}
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

type claveVuelo struct {
	fecha       string
	numeroVuelo int
}

func cmpClaveVuelo(a, b claveVuelo) int {
	fechaCmp := strings.Compare(a.fecha, b.fecha)
	if fechaCmp != 0 {
		return fechaCmp
	}
	aCadena := strconv.Itoa(a.numeroVuelo)
	bCadena := strconv.Itoa(b.numeroVuelo)
	return strings.Compare(aCadena, bCadena)
}

func CrearTablero() *TableroImpl {
	vuelosCodigo := diccionario.CrearHash[int, vuelo]()
	vuelosFecha := diccionario.CrearABB[claveVuelo, vuelo](cmpClaveVuelo)

	return &TableroImpl{vuelosCodigo: vuelosCodigo, vuelosFecha: vuelosFecha}
}

func procesarDatos(datos []string) vuelo {
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

func (tablero *TableroImpl) AgregarArchivo(archivo string) {
	file, err := os.Open(archivo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error en comando agregar_archivo")
		return
	}
	defer file.Close()
	s := bufio.NewScanner(file)
	for s.Scan() {
		linea := s.Text()
		datos := strings.Split(linea, ",")
		vuelo := procesarDatos(datos)
		clave := claveVuelo{fecha: vuelo.fecha, numeroVuelo: vuelo.numeroVuelo}
		if tablero.vuelosCodigo.Pertenece(vuelo.numeroVuelo) {
			info := tablero.vuelosCodigo.Obtener(vuelo.numeroVuelo)
			tablero.vuelosFecha.Borrar(claveVuelo{info.fecha, info.numeroVuelo})
			tablero.vuelosCodigo.Borrar(vuelo.numeroVuelo)
		}
		tablero.vuelosFecha.Guardar(clave, vuelo)
		tablero.vuelosCodigo.Guardar(vuelo.numeroVuelo, vuelo)
	}
}

func (tablero *TableroImpl) VerTablero(k int, modo string, desde string, hasta string) ([]string, error) {
	if k <= 0 || hasta < desde || (modo != "asc" && modo != "desc") {
		return nil, errors.New("Error en comando ver_tablero")
	}
	var resultado []string
	pilaAux := pila.CrearPilaDinamica[vuelo]()
	contador := 0

	for iter := tablero.vuelosFecha.IteradorRango(&claveVuelo{fecha: desde}, nil); iter.HaySiguiente() && contador < k; iter.Siguiente() {
		clave, vuelo := iter.VerActual()
		if clave.fecha > hasta {
			break
		}
		if modo == "desc" {
			pilaAux.Apilar(vuelo)
		} else {
			linea := fmt.Sprintf("%s - %d", vuelo.fecha, vuelo.numeroVuelo)
			resultado = append(resultado, linea)
			contador++
		}
	}
	if modo == "desc" {
		for !pilaAux.EstaVacia() && contador < k {
			valor := pilaAux.Desapilar()
			linea := fmt.Sprintf("%s - %d", valor.fecha, valor.numeroVuelo)
			resultado = append(resultado, linea)
			contador++
		}
	}
	return resultado, nil
}

func (tablero *TableroImpl) InfoVuelo(codigo int) error {
	if !tablero.vuelosCodigo.Pertenece(codigo) {
		return errors.New("Error en comando info_vuelo")
	}
	vuelo := tablero.vuelosCodigo.Obtener(codigo)
	fmt.Printf("%d %s %s %s %s %d %s %d %d %d\n",
		vuelo.numeroVuelo, vuelo.aerolinea, vuelo.origen, vuelo.destino,
		vuelo.matricula, vuelo.prioridad, vuelo.fecha, vuelo.atraso, vuelo.tiempoDeVuelo, vuelo.cancelado)
	return nil
}

func cmp(a, b vuelo) int {
	if a.prioridad > b.prioridad {
		return 1
	} else if a.prioridad < b.prioridad {
		return -1
	} else {
		aStr := strconv.Itoa(a.numeroVuelo)
		bStr := strconv.Itoa(b.numeroVuelo)
		if aStr < bStr {
			return 1
		} else if aStr > bStr {
			return -1
		} else {
			return 0
		}
	}
}

func TopK(arreglo []vuelo, k int) []vuelo {
	if k > len(arreglo) {
		k = len(arreglo)
	}
	heap := cola_prioridad.CrearHeapArr(arreglo, cmp)
	top := make([]vuelo, k)

	for i := 0; i < k; i++ {
		top[i] = heap.Desencolar()
	}
	return top
}
func (tablero *TableroImpl) PrioridadVuelos(k int) {
	arreglo := make([]vuelo, tablero.vuelosCodigo.Cantidad())
	i := 0
	for iter := tablero.vuelosCodigo.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelo := iter.VerActual()
		arreglo[i] = vuelo
		i++
	}
	topVuelos := TopK(arreglo, k)
	for _, elem := range topVuelos {
		fmt.Printf("%d - %d\n", elem.prioridad, elem.numeroVuelo)
	}
}

func (tablero *TableroImpl) SiguienteVuelo(origen, destino, fecha string) {
	for iter := tablero.vuelosFecha.IteradorRango(&claveVuelo{fecha: fecha}, nil); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelo := iter.VerActual()
		if vuelo.destino == destino && vuelo.origen == origen {
			tablero.InfoVuelo(vuelo.numeroVuelo)
			return
		}

	}
	fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fecha)
}

func (tablero *TableroImpl) Borrar(desde, hasta string) error {
	if desde > hasta {
		return errors.New("Error en comando borrar")
	}

	iter := tablero.vuelosFecha.IteradorRango(&claveVuelo{fecha: desde}, nil)
	for iter.HaySiguiente() {
		clave, vuelo := iter.VerActual()
		if clave.fecha > hasta {
			break
		}
		iter.Siguiente()
		tablero.vuelosFecha.Borrar(clave)
		tablero.InfoVuelo(vuelo.numeroVuelo)
		tablero.vuelosCodigo.Borrar(vuelo.numeroVuelo)
	}
	return nil
}
