package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const coferenceTickets = 75
	var remainingTickets = 50

	fmt.Println(conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets avaialble\n", remainingTickets)
	fmt.Printf("Get your tickets to attend the %v\n", conferenceName)

}
