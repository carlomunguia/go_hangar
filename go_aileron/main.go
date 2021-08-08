package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
		[
			{
				"first_name": "Clark",
				"last_name": "Kent",
				"hair_color": "black",
				"has_dog": true
			},
			{
				"first_name": "Bruce",
				"last_name": "Wayne",
				"hair_color": "black",
				"has_dog": false
			}

]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)

	var mySlice []Person

	var m1 Person

	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "Red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person

	m2.FirstName = "Dale"
	m2.LastName = "Cooper"
	m2.HairColor = "Black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "   ")
	if err != nil {
		log.Println("Error marshalling json", err)
	}

	fmt.Println(string(newJson))
}
