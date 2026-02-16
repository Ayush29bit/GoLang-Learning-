package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const coferenceTickets int = 75
	remainingTickets := 50

	fmt.Printf("conferenceTicket is %T, conferenceName is %T, remainingTickets is %T\n", coferenceTickets, conferenceName, remainingTickets)

	fmt.Println(conferenceName)
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets avaialble\n", remainingTickets)
	fmt.Printf("Get your tickets to attend the %v\n", conferenceName)

	var bookings []string

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets int

		// ask user for their details
		fmt.Println("Enter your First name")
		fmt.Scan(&firstName)
		fmt.Println("Enter yout last name")
		fmt.Scan(&lastName)
		fmt.Println("Enter your email")
		fmt.Scan(&email)

		// ask user for number of tickets
		fmt.Println("Enter number of tickets you want to buy")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		// Inserting user details into the array(slice for flexible size)
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice:%v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Print("Slice length: ", len(bookings), "\n")

		fmt.Printf("User %v %v with email %v booked %v tickets\n", firstName, lastName, email, userTickets)
		fmt.Println("Thank you for booking your tickets!")

		remainingTickets = remainingTickets - userTickets
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("The first names of bookings are: %v\n", firstNames)

		if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is booked out. Come back next year.")
			break
		}
	}

}
