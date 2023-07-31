package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {

	cnt_z := 0
	cnt_o := 0

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var target string
	status := -1
	fmt.Fscanln(reader, &target)
	for i := 0; i < len(target); i++ {
		temp := int(target[i])
		if temp == 48 {
			if status != temp {
				cnt_z += 1
			}
			status = 48
		} else {
			if status != temp {
				cnt_o += 1
			}
			status = 49
		}
	}
	fmt.Fprintln(writer, math.Min(float64(cnt_z), float64(cnt_o)))
}
