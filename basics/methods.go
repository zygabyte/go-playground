package main

import "fmt"

type user struct {
	name  string
	email string
}

func main() {
	james := user{"James White", "jameswhite@gmail.com"}
	ade := &user{"Ade Doe", "adedoe@gmail.com"}

	james.notifyUser()
	ade.notifyUser()

	james.changeEmail("jameswhite@yahoo.com")
	ade.changeEmail("jameswhite@yahoo.com")

	james.notifyUser()
	ade.notifyUser()
}

// u is a value receiver. since notifyUser has a value receiver, it's called a method.
// value receivers bind the calling of the function to the particular type that's passed.
func (u user) notifyUser() {
	fmt.Printf("sending email to %s with email %s\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	fmt.Printf("changing email from %s to %s\n", u.email, email)
	u.email = email
}
