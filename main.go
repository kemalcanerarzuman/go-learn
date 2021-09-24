package main

import "fmt"

var (
	helloWorld string
	name       string
	email      string
)

func main() {
	helloWorld = "Hello, World"
	fmt.Println(helloWorld)
	name, email = "Kemal Caner Arzuman", "kemalcanerarzuman@gmail.com"
	fmt.Printf("My name is %s, and my email is %s.", name, email)
}
