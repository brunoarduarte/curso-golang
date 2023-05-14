package main

import "time"

func main() {
	t := time.Now()
	switch { // switch true
	case t.Hour() < 12:
		println("Good morning!")
	case t.Hour() < 18:
		println("Good afternoon.")
	case t.Hour() < 22:
		println("Good evening.")
	default:
		println("Good night.")
	}
}
