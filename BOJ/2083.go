package main

import "fmt"

func main() {
	for true {
		var a string
		var b int
		var c int

		fmt.Scan(&a, &b, &c)

		if a == "#" && b == 0 && c == 0 {
			break
		}
		if b > 17 || c >= 80 {
			fmt.Println(a, "Senior")
		} else {
			fmt.Println(a, "Junior")
		}
	}
}
