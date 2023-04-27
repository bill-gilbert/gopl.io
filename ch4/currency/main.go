package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RUR
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RUR: "₽"}
	fmt.Println(RUR, symbol[RUR]) //  "3 ₽"
	r := [...]int{8: -1, -2}
	fmt.Printf("length %d", len(r))
}
