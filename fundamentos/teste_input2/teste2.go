package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	a, _ := strconv.ParseFloat(scanner.Text(), 64)
	x := 3.14159 * math.Pow(a, 2)
	fmt.Printf("A = %.4f\n", x)
}
