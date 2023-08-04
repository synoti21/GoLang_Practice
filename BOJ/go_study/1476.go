package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var e, s, m int
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscanln(reader, &e, &s, &m)

	se := 1
	ss := 1
	sm := 1

	for i := 1; i <= 7980; i++ {
		if se == e && ss == s && sm == m {
			fmt.Fprintln(writer, i)
		}
		se += 1
		ss += 1
		sm += 1
		if se > 15 {
			se = 1
		}
		if ss > 28 {
			ss = 1
		}
		if sm > 19 {
			sm = 1
		}
	}
}
