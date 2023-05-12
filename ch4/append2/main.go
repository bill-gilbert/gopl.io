// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 88.

// Append illustrates the behavior of the built-in append function.
package main

import "fmt"

func appendslice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to expand the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

// !+append
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	fmt.Printf("len=%d cap=%d \n", len(x), cap(x))
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

//!-append

// !+growth
func main() {
	var x []int
	x = appendslice(x, 1)
	x = appendslice(x, 2, 3)
	x = appendslice(x, 4, 5, 6)
	x = appendslice(x, x...) // Добавление среза x
	fmt.Println(x)           //  "[1 2 3 4 5 6 1 2  3 4 5 6]"
}
