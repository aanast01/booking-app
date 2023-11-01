package helper

import (
	"fmt"
	"strings"
)

func GreetUser(conferenceName string, remainingTickets uint) {
	// Print welcome message
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v remaining tickets\n", remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func GetUserInput(bookings_names *[]string, bookings_tickets *[]uint, remainingTickets *uint) string {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	var isValidName = false
	for !isValidName {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)
		isValidName = len(firstName) >= 2
		if !isValidName {
			fmt.Printf("Invalid first name. First name must be at least 2 characters\n")
		}
	}

	isValidName = false
	for !isValidName {
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)
		isValidName = len(lastName) >= 2
		if !isValidName {
			fmt.Printf("Invalid last name. Last name must be at least 2 characters\n")
		}
	}

	var isValidEmail = false
	for !isValidEmail {
		fmt.Println("Enter your email: ")
		fmt.Scan(&email)
		isValidEmail = strings.Contains(email, "@")
		if !isValidEmail {
			fmt.Printf("Invalid email. Email must contain the @ character\n")
		}
	}

	var isValidTicketNumber = false
	for !isValidTicketNumber {
		fmt.Println("Enter how many tickets you want: ")
		fmt.Scan(&userTickets)
		isValidTicketNumber = userTickets > 0 && userTickets <= *remainingTickets
		if !isValidTicketNumber {
			fmt.Printf("Invalid number of tickets. Ticket number must be between 1 and %v\n", *remainingTickets)
		}
	}

	*bookings_names = append(*bookings_names, firstName+" "+lastName)
	*bookings_tickets = append(*bookings_tickets, userTickets)
	*remainingTickets = *remainingTickets - userTickets

	return email
}

func PrintReservations(bookings_names *[]string, bookings_tickets *[]uint) {
	fmt.Printf("----- All Reservations -----\n")

	var tickets = *bookings_tickets
	for index, booking := range *bookings_names {
		fmt.Printf("Reservation %v:\nName: %v    Tickets: %v\n", index+1, strings.Fields(booking)[0], tickets[index])
	}
}
