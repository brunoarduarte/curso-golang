package main

import "fmt"

func imprimirAprovados(aprovados ...string) {
	println("Lista de Aprovados")
	for i, aprovado := range aprovados {
		fmt.Printf("%d) %s\n", i+1, aprovado)
	}
}

func main() {
	aprovados := []string{"Maria", "Pedro", "Rafael", "Guilherme"}
	imprimirAprovados(aprovados...) // slice como parametro. Funciona com slice e não com array
}
