package main

import (
	"fmt"
	"strings"
)

// Defining Package variables
var conferenceName string = "Go conference"

const coferenceTickets int = 75

var remainingTickets int = 50
var bookings = []string{}

func main() {

	greetUsers()

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

		// validate user input
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			fmt.Println("Your input is valid, processing your booking...")

			// Inserting user details into the array(slice for flexible size)

			bookings = bookTickets(firstName, lastName, userTickets)

			fmt.Printf("User %v %v with email %v booked %v tickets\n", firstName, lastName, email, userTickets)
			fmt.Println("Thank you for booking your tickets!")

			var firstNames = getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

		} else {
			if !isValidName {
				fmt.Println("First name and last name should be at least 2 characters long")
			}
			if !isValidEmail {
				fmt.Println("Email address should contain @ sign")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets should be more than 0 and less than remaining tickets")
			}
			fmt.Println("Your input is invalid, please try again")
			continue
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets avaialble\n", remainingTickets)
	fmt.Printf("Get your tickets to attend the %v\n", conferenceName)
}
func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets int) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets
	return isValidName, isValidEmail, isValidTicketNumber
}

func bookTickets(firstName string, lastName string, userTickets int) []string {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)
	return bookings

}
