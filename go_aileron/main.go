package main

import (
	"log"
	"sort"
)

func main() {
	var miSlice []string

	miSlice = append(miSlice, "Pink")
	miSlice = append(miSlice, "Moon")

	log.Println(miSlice)

	numbers := []int{9, 1, 3, 6, 2, 3, 1, 3, 2}
	sort.Ints(numbers)
	log.Println(numbers)

	var isTrue bool
	var isFalse bool

	isTrue = true
	isFalse = false

	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isFalse is", isFalse)
	}

	cat := "cat"

	if cat == "cat" {
		log.Println("meow meow")
	}

	myNum := 100

	if myNum > 99 && isTrue {
		log.Println("totally normal conditional stuff")
	}
}
