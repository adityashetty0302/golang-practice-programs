package main

import "fmt"

type bot interface {
	getGreetings() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreetings(eb)
	printGreetings(sb)

}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

func (e englishBot) getGreetings() string {
	return "Hello"
}

func (e spanishBot) getGreetings() string {
	return "Hola"
}
