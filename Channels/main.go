package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://www.google.com",
		"http://www.facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//? communicating through the channel with values typed string
	c := make(chan string)

	for _, link := range links {
		//! starting this function in a new go routine using go.  Because there is a main routine and then we are creating child routines when the
		//! main routine finishes while the child routines are still doing work the program quits before finishing the child routines work
		//! so if we just add the go in front of the items it will not print because it closes before getting to that part of the functions
		//! in the child routines so we need to use a channel to help the main routine watch for when the child routines finish their code
		//! channels are the only way for go routines to communicate with each other.  Channels are typed the same way as any other value such as
		//! a struct or any other type.  So the channel will only be able to send that type from that channel.
		go checkLink(link, c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

//? ****** Ways of sending commmunication within channels *******
//? channel <- 5    		=== Send the value '5' into this channel
//? myNumber <- channel 	=== Wait for a value to be sent into the channel.  When we get one assign the value to 'myNumber'
//? fmt.Println(<- channel) === Wait for a value to be sent into the channel.  When we get one, log it out immediately

// ? Taking in the channel as an argument for the function that way we can pass the channel in from the main function needs typed with the chan for
// ? channel and the type of the channel data that can be shared
func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + " errored and might be down!")
		c <- "Might be down!"
		return
	}

	fmt.Println(link + " appears to be working!")
	c <- "It's appears to be working!"
}
