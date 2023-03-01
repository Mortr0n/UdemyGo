package main

import "fmt"

func main() {
	myInts := []int{0}

	i := 1
	for i <= 10 {
		myInts = append(myInts, i)
		i++
	}

	for _, myInt := range myInts {
		if myInt%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}

	}

	fmt.Printf("myInts: %v", myInts)

}
