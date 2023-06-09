package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	// fmt.Println(x / y)
	fmt.Println(x / float64(y))

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println(notaFinal)

	// cuidado na hora de concatenar string
	fmt.Println("Teste " + string(rune(97)))

	// int para string
	fmt.Println("Teste " + strconv.Itoa(123))

	// string para int
	num, _ := strconv.Atoi("123")
	fmt.Println(num - 122)

	// string to bool
	b, _ := strconv.ParseBool("true")
	if b {
		fmt.Println("Verdadeiro")
	}
}
