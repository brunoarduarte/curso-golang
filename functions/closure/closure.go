package main

func closure() func() {
	x := 10
	var funcao = func() {
		println(x)
	}
	return funcao
}

func main() {
	x := 20
	println(x)

	imprimeX := closure()
	imprimeX()
}
