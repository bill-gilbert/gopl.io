// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 88.

// Append illustrates the behavior of the built-in append function.
package main

import (
	"fmt"
	"sort"
)

//func remove(slice []int, i int) []int {
//	//slice[i] = slice[len(slice)1]
//	return slice[:len(slice)1]
//}

func remove(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	fmt.Println(i)
	return slice[:len(slice)-1]
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
		//
		// if xv != y[k] {
		// 	return false
		// }
	}
	return true
}

func main() {
	s := []int{5, 6, 7, 8, 9}
	fmt.Println(remove(s, 2))

	ages := map[string]int{"hello": 1, "pavel": 2, "apple": 33}
	fmt.Println(ages["hello"], " ", ages["pavel"])
	fmt.Println(ages["superhero"] + 1)

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	fmt.Println("-------------------------------")
	var names []string
	for name := range ages {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
	fmt.Println("-------------------------------")
	var test map[string]int
	fmt.Printf("length is %d", len(test))

	test = make(map[string]int, 10)
	test["test"] = 12
	fmt.Printf("element test has value %d \n", test["test"])

	var age, res = test["unknown"]

	if !res {
		fmt.Printf("unknown %t is err \n", res)
	} else {
		fmt.Printf("unknown %d \n", age)
	}

	//	firstMap := map[string]int{"One": 1, "Two": 2}
	//	secondMap := map[string]int{"One": 1, "Two": 2}
	//	fmt.Println(equal(firstMap, secondMap))
	fmt.Println(equal(map[string]int{"A": 0}, map[string]int{"B": 42}))
}
