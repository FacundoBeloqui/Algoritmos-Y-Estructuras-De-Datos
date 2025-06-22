package algueiza

type Tablero interface {
	//AgregarVuelo inserta un nuevo vuelo o actualiza los datos en caso de que ya exista el codigo
	AgregarVuelo(vuelo)

	//VerTablero devuelve los primeros k vuelos de fentre las fechas desde y hasta, inclusive, ordenados por fecha en modo "asc" o "desc". Ademas, devuelve nil o un mensaje de error (en caso de que k sea menor o igual a cero, hasta sea menor a desde o el modo no sea "asc" o "desc")
	VerTablero(int, string, string, string) ([][]string, error)

	//InfoVuelo devuelve la informacion completa de un vuelo dado su codigo y devuelve nil o un mensaje de error (en caso de que no exista el codigo de vuelo)
	InfoVuelo(int) ([]string, error)

	//PrioridadVuelos retorna los primeros k vuelos con mayor prioridad
	PrioridadVuelos(int) []string

	//SiguienteVuelo retorna todos los vuelos de origen y destino dado desde la fecha dada, inclusive. Ademas, devuelve true(en caso de que se encuentre por lo menos un vuelo) o nil y false(en caso de que no se encuentre ningun vuelo)
	SiguienteVuelo(string, string, string) ([]string, bool)

	//Borrar elimina todos los vuelos del tablero cuya fecha esté entre las fechas dadas y devuelve estos vuelos. Ademas, devuelve nil o un mensaje de error(en caso de que desde sea mayor a hasta)
	Borrar(string, string) ([]string, error)
}
