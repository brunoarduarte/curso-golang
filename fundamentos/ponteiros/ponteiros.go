package main

import "fmt"

func main() {
	i := 1

	var p *int = nil
	
	p = &i // pegando o endereço da variável
	*p++   // desreferenciando (pegando o valor)
	i++
	
	// O GO não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
