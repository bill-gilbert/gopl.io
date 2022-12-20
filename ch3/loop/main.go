package main

import "fmt"

func main() {
	medals := []string{"золото", "серебро", "бронза"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) //  "бронза",  "серебро",  "золото”
	}

}
