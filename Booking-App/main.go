package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const coferenceTickets = 75
	var remainingTickets = 50

	fmt.Println(conferenceName)
	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have", remainingTickets, "tickets avaialble")
	fmt.Println("Get your tickets to attend the", conferenceName)

}
