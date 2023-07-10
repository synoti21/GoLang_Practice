package main

import (
	"fmt"
	"math"
)

func add(a int, b int, c int) int {
	if a == b && b == c {
		return 10000 + (a * 1000)
	} else if a == b {
		return 1000 + (a * 100)
	} else if b == c {
		return 1000 + (b * 100)
	} else if a == c {
		return 1000 + (a * 100)
	} else {
		return int(math.Max(math.Max(float64(a), float64(b)), float64(c))) * 100
	}
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	fmt.Print(add(a, b, c))
}
