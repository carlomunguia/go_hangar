package main

import "log"

type miStruct struct {
	FirstName string
}

func (m *miStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var miVar miStruct
	miVar.FirstName = "John"

	miVar2 := miStruct{
		FirstName: "Mary",
	}

	log.Println("miVar is set to:", miVar.printFirstName())
	log.Println("mivar2 is set to: ", miVar2.printFirstName())
}
