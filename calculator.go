package main

import (
	"fmt"

	"github.com/shimizu/calculator/internal/add"
	"github.com/shimizu/calculator/internal/subtract"
)

func main() {
	a, b := 10, 5
	sum, difference := calculate(a, b)
	fmt.Printf("%d + %d = %d\n", a, b, sum)
	fmt.Printf("%d - %d = %d\n", a, b, difference)
}

func calculate(a, b int) (int, int) {
	sum := add.Add(a, b)
	difference := subtract.Subtract(a, b)
	return sum, difference
}
