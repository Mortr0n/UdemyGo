package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00FF00",
	// 	"blue":  "#0000FF",
	// }

	colors["white"] = "#FFFFFF"
	colors["blue"] = "#0000FF"
	colors["green"] = "00FF00"
	colors["yellow"] = "#FFFF00"
	colors["red"] = "#FF0000"

	// delete(colors, "white")

	// for i, val := range colors {
	// 	//! did this as a test. this won't print keys.  See below function to print the keys.
	// 	fmt.Println("I is ", i, " val is : ", val)
	// }

	fmt.Println(colors["white"])

	printMap(colors)
}

func printMap(c map[string]string) {
	//? color and hex can be thought of as key and value and often we'll probably just use k,v
	//? for k, v := range c
	for color, hex := range c {
		fmt.Println("Color is :", color, " - hex is :", hex)
	}
}
