package main

import "fmt"

func main() {
	type user struct {
		name         string
		email        string
		phone_number int
		registered   bool
	}
	type admin struct {
		user  user
		level string
	}

	fred := user{"Fred Smith", "fredsmith@gmail.com", 123455, false}
	fmt.Println(fred)

	james := admin{fred, "1000"}
	fmt.Println(james)
	fmt.Println(james.user.email)

}
