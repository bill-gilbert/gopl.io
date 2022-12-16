package main

import (
	"fmt"
	"gopl.io/ch2/popcount"
)

func main() {
	i := newInt2()
	*i = 2
	fmt.Println(*i)

	j := newInt2()
	*j = 4
	fmt.Println(*j)
	fmt.Println(popcount.PopCount(0x1234567890ABCDEF))
}

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}
