package main

import (
	"fmt"
	"time"
)

func tipo(i interface{}) string {
	switch i.(type) {
	case int:
		return "int"
	case float32, float64:
		return "float"
	case string:
		return "string"
	case func():
		return "func"
	case bool:
		return "bool"
	default:
		return "não sei"
	}
}

func main() {
	fmt.Println(tipo(2.3))
	fmt.Println(tipo(1))
	fmt.Println(tipo("Opa"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(time.Now()))
	fmt.Println(tipo(true))
}
