package main

func obterValorAprovado(numero int) int {
	defer println("Fim!")
	if numero > 5000 {
		println("Valor alto...")
		return 5000
	}
	println("Valor baixo...")
	return numero
}

func main() {
	println(obterValorAprovado(6000))
	println(obterValorAprovado(3000))
}
