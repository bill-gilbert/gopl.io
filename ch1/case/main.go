// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
)

func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		switch i {
		case 1:
			fmt.Println("нечетные")
		case 3:
			fmt.Println("нечетные")
		case 5:
			fmt.Println("нечетные")
		case 2:
			fmt.Println("чётные")
		case 4:
			fmt.Println("чётные")
		default:
			fmt.Println("что-то другое")
		}
	}
}

//!-
