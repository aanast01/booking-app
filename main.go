package main

import (
	"booking-app/helper"
	"fmt"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	var bookings_names []string
	var bookings_tickets []uint

	for remainingTickets > 0 {
		// greet user
		helper.GreetUser(conferenceName, remainingTickets)

		// ask user for their name
		var email = helper.GetUserInput(&bookings_names, &bookings_tickets, &remainingTickets)

		fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v.\n", bookings_names[len(bookings_names)-1], bookings_tickets[len(bookings_tickets)-1], email)
		fmt.Printf("Remaining tickets for %v: %v\n\n", conferenceName, remainingTickets)
	}

	fmt.Printf("Conference %v is fully booked. Thank you!!!\n\n", conferenceName)

	// print all reservations
	helper.PrintReservations(&bookings_names, &bookings_tickets)
}
