package algueiza

import (
	"errors"
	"tdas/cola_prioridad"
	"tdas/pila"

	"fmt"
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

func (tablero *TableroImpl) AgregarVuelo(vuelo vuelo) {
	clave := claveVuelo{fecha: vuelo.fecha, numeroVuelo: vuelo.numeroVuelo}

	if tablero.vuelosCodigo.Pertenece(vuelo.numeroVuelo) {
		info := tablero.vuelosCodigo.Obtener(vuelo.numeroVuelo)
		tablero.vuelosFecha.Borrar(claveVuelo{info.fecha, info.numeroVuelo})
		tablero.vuelosCodigo.Borrar(vuelo.numeroVuelo)
	}
	tablero.vuelosFecha.Guardar(clave, vuelo)
	tablero.vuelosCodigo.Guardar(vuelo.numeroVuelo, vuelo)

}

func (tablero *TableroImpl) VerTablero(k int, modo string, desde string, hasta string) ([][]string, error) {
	if k <= 0 || hasta < desde || (modo != "asc" && modo != "desc") {
		return nil, errors.New("Error en comando ver_tablero")
	}
	var resultado [][]string
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
			linea := []string{vuelo.fecha, strconv.Itoa(vuelo.numeroVuelo)}
			resultado = append(resultado, linea)
			contador++
		}
	}
	if modo == "desc" {
		for !pilaAux.EstaVacia() && contador < k {
			valor := pilaAux.Desapilar()
			linea := []string{valor.fecha, strconv.Itoa(valor.numeroVuelo)}
			resultado = append(resultado, linea)
			contador++
		}
	}
	return resultado, nil
}

func (tablero *TableroImpl) InfoVuelo(codigo int) ([]string, error) {
	if !tablero.vuelosCodigo.Pertenece(codigo) {
		return nil, errors.New("Error en comando info_vuelo")
	}
	vuelo := tablero.vuelosCodigo.Obtener(codigo)

	return []string{
		strconv.Itoa(vuelo.numeroVuelo),
		vuelo.aerolinea,
		vuelo.origen,
		vuelo.destino,
		vuelo.matricula,
		strconv.Itoa(vuelo.prioridad),
		vuelo.fecha,
		strconv.Itoa(vuelo.atraso),
		strconv.Itoa(vuelo.tiempoDeVuelo),
		strconv.Itoa(vuelo.cancelado),
	}, nil
}

func (vuelo vuelo) Cmp(otroVuelo vuelo) int {
	if vuelo.prioridad > otroVuelo.prioridad {
		return 1
	} else if vuelo.prioridad < otroVuelo.prioridad {
		return -1
	} else {
		vueloStr := strconv.Itoa(vuelo.numeroVuelo)
		otroVueloStr := strconv.Itoa(otroVuelo.numeroVuelo)
		if vueloStr < otroVueloStr {
			return 1
		} else if vueloStr > otroVueloStr {
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
	heap := cola_prioridad.CrearHeapArr(arreglo, vuelo.Cmp)
	top := make([]vuelo, k)

	for i := 0; i < k; i++ {
		top[i] = heap.Desencolar()
	}
	return top
}
func (tablero *TableroImpl) PrioridadVuelos(k int) []string {
	arreglo := make([]vuelo, tablero.vuelosCodigo.Cantidad())
	i := 0
	for iter := tablero.vuelosCodigo.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelo := iter.VerActual()
		arreglo[i] = vuelo
		i++
	}
	topVuelos := TopK(arreglo, k)

	var resultado []string
	for _, elem := range topVuelos {
		resultado = append(resultado, fmt.Sprintf("%d - %d", elem.prioridad, elem.numeroVuelo))
	}
	return resultado
}

func (tablero *TableroImpl) SiguienteVuelo(origen, destino, fecha string) ([]string, bool) {
	encontrado := false
	for iter := tablero.vuelosFecha.IteradorRango(&claveVuelo{fecha: fecha}, nil); iter.HaySiguiente(); iter.Siguiente() {
		_, vuelo := iter.VerActual()
		if vuelo.destino == destino && vuelo.origen == origen {
			encontrado = true
			info, _ := tablero.InfoVuelo(vuelo.numeroVuelo)
			return info, encontrado
		}
	}
	return nil, encontrado
}

func (tablero *TableroImpl) Borrar(desde, hasta string) ([]string, error) {
	if desde > hasta {
		return nil, errors.New("Error en comando borrar")
	}
	var resultados []string
	iter := tablero.vuelosFecha.IteradorRango(&claveVuelo{fecha: desde}, nil)

	for iter.HaySiguiente() {
		clave, vuelo := iter.VerActual()
		if clave.fecha > hasta {
			break
		}

		info, _ := tablero.InfoVuelo(vuelo.numeroVuelo)
		resultados = append(resultados, strings.Join(info, " "))

		iter.Siguiente()
		tablero.vuelosFecha.Borrar(clave)
		tablero.vuelosCodigo.Borrar(vuelo.numeroVuelo)
	}

	return resultados, nil
}
