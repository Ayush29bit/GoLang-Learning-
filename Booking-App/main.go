package main

import "fmt"

func main() {
	conferenceName := "Go conference"
	const coferenceTickets int = 75
	remainingTickets := 50

	fmt.Printf("conferenceTicket is %T, conferenceName is %T, remainingTickets is %T\n", coferenceTickets, conferenceName, remainingTickets)

	fmt.Println(conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets avaialble\n", remainingTickets)
	fmt.Printf("Get your tickets to attend the %v\n", conferenceName)

	var userName string
	var userTickets int
	// ask user for their name
	fmt.Println("Enter your name")
	fmt.Scan(userName)
	// ask user for number of tickets
	fmt.Println("Enter number of tickets you want to buy")
	fmt.Scan(&userTickets)

	fmt.Printf("Use r %v booked %v tickets\n", userName, userTickets)
}
