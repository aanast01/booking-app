package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	var bookings_names []string
	var bookings_tickets []uint

	for remainingTickets > 0 {
		// Print welcome message
		fmt.Printf("Welcome to %v booking application\n", conferenceName)
		fmt.Printf("We have a total of %v tickets and %v are available\n", conferenceTickets, remainingTickets)
		fmt.Println("Get your tickets here to attend")

		var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter how many tickets you want: ")
		fmt.Scan(&userTickets)

		if userTickets > remainingTickets {
			fmt.Printf("Sorry there only %v available tickets\n\n", remainingTickets)
			continue
		}

		bookings_names = append(bookings_names, firstName+" "+lastName)
		bookings_tickets = append(bookings_tickets, userTickets)
		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v.\n", bookings_names[len(bookings_names)-1], bookings_tickets[len(bookings_tickets)-1], email)
		fmt.Printf("Remaining tickets for %v: %v\n\n", conferenceName, remainingTickets)
	}

	fmt.Printf("Conference %v is fully booked. Thank you!!!\n\n", conferenceName)

	fmt.Printf("----- All Reservations -----\n")

	for index, booking := range bookings_names {
		fmt.Printf("Reservation %v:\nName: %v    Tickets: %v\n", index+1, strings.Fields(booking)[0], bookings_tickets[index])
	}
}
