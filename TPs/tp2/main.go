package main

import (
	"bufio"
	"fmt"
	"os"
	"tp2/algueiza"
)

func main() {
	tablero := algueiza.CrearTablero()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linea := scanner.Text()
		if linea == "" {
			break
		}

		err := algueiza.ProcesarComando(linea, tablero)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		} else {
			fmt.Println("OK")
		}
	}
}
