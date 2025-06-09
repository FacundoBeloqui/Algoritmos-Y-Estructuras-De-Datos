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
	return vuelo{numeroVuelo: numeroVuelo, aerolinea: datos[1], origen: datos[2], destino: datos[3], matricula: datos[4], prioridad: prioridad, fecha: fecha, atraso: atraso, tiempoDeVuelo: tiempoDeVuelo, cancelado: cancelado}
}
func (t *TableroImpl) AgregarArchivo(archivo string) {
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
		datosVuelo := procesarDatos(datos)
		clave := claveVuelo{fecha: datosVuelo.fecha, numeroVuelo: datosVuelo.numeroVuelo}
		if t.vuelosCodigo.Pertenece(datosVuelo.numeroVuelo) {
			info := t.vuelosCodigo.Obtener(datosVuelo.numeroVuelo)
			t.vuelosFecha.Borrar(claveVuelo{info.fecha, info.numeroVuelo})
			t.vuelosCodigo.Borrar(datosVuelo.numeroVuelo)
		}
		t.vuelosFecha.Guardar(clave, datosVuelo)
		t.vuelosCodigo.Guardar(datosVuelo.numeroVuelo, datosVuelo)
	}
}

func (t *TableroImpl) VerTablero(k int, modo string, desde string, hasta string) {
	if k <= 0 || hasta < desde || (modo != "asc" && modo != "desc"){
		fmt.Fprintf(os.Stderr, "Error en comando ver_tablero\n")
		return
	}
	pilaAux := pila.CrearPilaDinamica[vuelo]()
	contador := 0
	for iter := t.vuelosFecha.IteradorRango(&claveVuelo{fecha: desde}, &claveVuelo{fecha: hasta}); iter.HaySiguiente() && contador < k; iter.Siguiente() {
		_, datos := iter.VerActual()
		if modo == "desc" {
			pilaAux.Apilar(datos)
		} else {
			fmt.Printf("%s - %d\n", datos.fecha, datos.numeroVuelo)
			contador++
		}
		
	}
	for !pilaAux.EstaVacia() && contador < k {
		valor := pilaAux.Desapilar()
		fmt.Printf("%s - %d\n", valor.fecha, valor.numeroVuelo)
		contador++
	}
}

func (t *TableroImpl) InfoVuelo(codigo int) error {
	if !t.vuelosCodigo.Pertenece(codigo) {
		return errors.New("Error en comando info_vuelo")
	}
	datos := t.vuelosCodigo.Obtener(codigo)
	fmt.Printf("%d %s %s %s %s %d %s %d %d %d\n",
		datos.numeroVuelo, datos.aerolinea, datos.origen, datos.destino,
		datos.matricula, datos.prioridad, datos.fecha, datos.atraso, datos.tiempoDeVuelo, datos.cancelado)
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

func TopK(arr []vuelo, k int) []vuelo {
	if k > len(arr){
		k = len(arr)
	}
	cp := cola_prioridad.CrearHeapArr(arr, cmp)
	top := make([]vuelo, k)

	for i := 0; i < k; i++ {
		top[i] = cp.Desencolar()
	}
	return top
}
func (t *TableroImpl) PrioridadVuelos(k int) {
	miarr := make([]vuelo, t.vuelosCodigo.Cantidad())
	i := 0
	for iter := t.vuelosCodigo.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		_, datos := iter.VerActual()
		miarr[i] = datos
		i++
	}
	topVuelos := TopK(miarr, k)
	for _, elem := range topVuelos {
		fmt.Printf("%d - %d\n", elem.prioridad, elem.numeroVuelo)
	}
}

func (t *TableroImpl) SiguienteVuelo(origen, destino, fecha string) {
	for iter := t.vuelosFecha.IteradorRango(&claveVuelo{fecha: fecha}, nil); iter.HaySiguiente(); iter.Siguiente() {
		_, datos := iter.VerActual()
		if datos.destino == destino && datos.origen == origen {
			t.InfoVuelo(datos.numeroVuelo)
			return
		}
		
	}
	fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fecha)
}

func (t *TableroImpl) Borrar(desde, hasta string) {
	if desde > hasta {
		fmt.Fprintf(os.Stderr, "Error en comando borrar\n")
		return
	}
	iter := t.vuelosFecha.IteradorRango(&claveVuelo{fecha: desde}, &claveVuelo{fecha: hasta})
	for iter.HaySiguiente(){
		clave, datos := iter.VerActual()
		iter.Siguiente()
		t.vuelosFecha.Borrar(clave)
		t.InfoVuelo(datos.numeroVuelo)
		t.vuelosCodigo.Borrar(datos.numeroVuelo)
	}
}