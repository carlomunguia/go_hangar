package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gF or %gC\n", f, c)
}

func main2() {
	const freezingF, boilingF = 32.0, 212.0
	fmt.Printf("%gF = %gC\n", freezingF, fToC(freezingF))
	fmt.Printf("%gF = %gC\n", boilingF, fToC(boilingF))
}

	func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}

v := 1
incr(&v)
fmt.PrintLn(incr(&v))

