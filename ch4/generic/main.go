// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 88.

// Append illustrates the behavior of the built-in append function.
package main

import (
	"encoding/json"
)

/*
*

	func PrintSlice[T any](s []T) {
	    for _, v := range s{
	        print(v)
	    }
	}
*/
func sum[T any](v any, input T) T {
	b, _ := json.Marshal(v)
	json.Unmarshal(b, &input)
	return input
}

// !+growth
func main() {

}
