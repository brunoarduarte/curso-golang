package main

import "fmt"

func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		println(n)
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
