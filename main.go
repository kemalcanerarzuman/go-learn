package main

import "fmt"

var (
	helloWorld string
	user       Person
)

type Person struct {
	name, email string
	age         int
}

func main() {
	helloWorld = "Hello, World"
	fmt.Println(helloWorld)
	user.name, user.email, user.age = "Kemal Caner Arzuman", "kemalcanerarzuman@gmail.com", 29
	fmt.Printf("User's name is %s, and email is %s, and age is %d.", user.name, user.email, user.age)
}
