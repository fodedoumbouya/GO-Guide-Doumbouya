package main

import "fmt"

// We are creating our Structure
type Person struct {
	Name string
	Age  int
	Sexe string
}

//https://golangdocs.com/channels-in-golang
// this Function only help us to put value in the channek
func SendToChannel(channel chan Person, person Person) {
	channel <- person
}

func main() {

	// Creating  empty channel
	channel := make(chan Person)
	// creating our Person
	p := Person{
		Name: "name",
		Age:  20,
		Sexe: "Female",
	}

	// calling our function `SendToChannel` by passing our channel and the personne
	// but we are executing it in goroutine that's why we have 'go' before our function
	go SendToChannel(channel, p)

	// we are saying here to wait for the out and then give the name to our varible 'name'
	name := (<-channel).Name
	// Display the name
	fmt.Println(name)

}
