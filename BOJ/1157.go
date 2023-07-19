package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var target string
	var freq = map[int]int{}
	var dup bool = false
	var maxv int = -1
	var maxk int = 0

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 97; i <= 122; i++ {
		freq[i] = 0
	}

	fmt.Fscanln(reader, &target)

	for i := 0; i < len(target); i++ {
		temp := int(target[i])
		if temp < 97 {
			temp += 32
		}
		freq[temp] += 1
	}

	for i := 97; i <= 122; i++ {
		if maxv == freq[i] {
			dup = true
		} else if freq[i] > maxv {
			dup = false
			maxv = freq[i]
			maxk = i
		}
	}

	if dup {
		fmt.Fprintln(writer, "?")
	} else {
		fmt.Fprintln(writer, string(maxk-32))
	}
}
