package main

import "fmt"

//http://goinbigdata.com/golang-pass-by-pointer-vs-pass-by-value/
//Pass by pointer vs pass by value in Go

type Person struct {
	firstName string
	lastName  string
}

func changeName(p Person) {
	p.firstName = "Bob"
}

func changePointerName(p *Person) {
	p.firstName = "Bob"
}

func main() {
	/*person := Person {
		firstName: "Alice",
		lastName: "Dow",
	}*/

	person := Person{"Alice", "Dow"}

	changeName(person)
	fmt.Println(person)

	changePointerName(&person)
	fmt.Println(person)
}
