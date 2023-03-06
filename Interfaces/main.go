package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sp spanishBot) {
// 	fmt.Println(sp.getGreeting())
// }

func (eb englishBot) getGreeting() string {
	// Super incredibly custom logic for generating an english greeting
	return ("Hello Nurse!!!")
}

// Can omit the value for the receiver object that we take in if we aren't making use of it inside the functioin
// See how eb is included in the englishBot greeting above and no sp for the spanishBot here.
func (spanishBot) getGreeting() string {
	return ("Hola Seniorita!")
}
