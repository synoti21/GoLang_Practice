package main

import "fmt"

func sol() {
	var n int
	var sum, input int64
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&input)
		sum += input
	}
	if sum == 0 {
		fmt.Println(0)
	} else if sum > 0 {
		fmt.Println("+")
	} else {
		fmt.Println("-")
	}
}
func main() {
	for i := 0; i < 3; i++ {
		sol()
	}
}
