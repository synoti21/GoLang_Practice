package main

import (
	"fmt"
	"math"
)

var color = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

func main() {
	var res [3]string
	var ans int

	for i := 0; i < 3; i++ {
		fmt.Scanln(&res[i])
	}
	ans = (color[res[0]]*10 + color[res[1]]) * int(math.Pow(10, float64(color[res[2]])))

	fmt.Println(ans)
}
