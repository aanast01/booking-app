package main

import (
	"booking-app/helper"
	"fmt"
)

func main() {
	var conferenceName string = "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = conferenceTickets
	var bookings []helper.Booking

	for remainingTickets > 0 {
		// greet user
		helper.GreetUser(conferenceName, remainingTickets)

		// ask user for their name
		helper.BookTicket(&bookings, &remainingTickets)
	}

	fmt.Printf("Conference %v is fully booked. Thank you!!!\n\n", conferenceName)

	// print all reservations
	helper.PrintReservations(&bookings)

	helper.WG.Wait()
}
