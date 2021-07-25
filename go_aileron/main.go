package main

import "log"

func main() {
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Carlo", "Smith", "carlo@smith.com", 32})
	users = append(users, User{"Leon", "Smith", "leon@smith.com", 22})
	users = append(users, User{"Paco", "Smith", "paco@smith.com", 29})

	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Age, l.Email)
	}
}
