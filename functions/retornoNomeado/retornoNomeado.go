package main

func trocar(p1, p2 int) (segundo, primeiro int) {
	segundo = p2
	primeiro = p1
	return //retorno limpo
}

func main() {
	x, y := trocar(2, 3)
	println(x, y)

	segundo, primeiro := trocar(7, 1)
	println(segundo, primeiro)
}
