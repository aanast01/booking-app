package helper

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var WG = sync.WaitGroup{}

type Booking struct {
	firstName string
	lastName  string
	email     string
	tickets   uint
}

func GreetUser(conferenceName string, remainingTickets uint) {
	// Print welcome message
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v remaining tickets\n", remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func BookTicket(bookings *[]Booking, remainingTickets *uint) {

	var userData Booking

	var isValidName = false
	for !isValidName {
		fmt.Println("Enter your first name: ")
		fmt.Scan(&userData.firstName)
		isValidName = len(userData.firstName) >= 2
		if !isValidName {
			fmt.Printf("Invalid first name. First name must be at least 2 characters\n")
		}
	}

	isValidName = false
	for !isValidName {
		fmt.Println("Enter your last name: ")
		fmt.Scan(&userData.lastName)
		isValidName = len(userData.lastName) >= 2
		if !isValidName {
			fmt.Printf("Invalid last name. Last name must be at least 2 characters\n")
		}
	}

	var isValidEmail = false
	for !isValidEmail {
		fmt.Println("Enter your email: ")
		fmt.Scan(&userData.email)
		isValidEmail = strings.Contains(userData.email, "@")
		if !isValidEmail {
			fmt.Printf("Invalid email. Email must contain the @ character\n")
		}
	}

	var isValidTicketNumber = false
	for !isValidTicketNumber {
		fmt.Println("Enter how many tickets you want: ")
		fmt.Scan(&userData.tickets)
		isValidTicketNumber = userData.tickets > 0 && userData.tickets <= *remainingTickets
		if !isValidTicketNumber {
			fmt.Printf("Invalid number of tickets. Ticket number must be between 1 and %v\n", *remainingTickets)
		}
	}

	*bookings = append(*bookings, userData)
	*remainingTickets = *remainingTickets - userData.tickets

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v.\n\n", userData.firstName, userData.tickets, userData.email)

	WG.Add(1)
	go sendTicket(userData)
}

func PrintReservations(bookings *[]Booking) {
	fmt.Printf("----- All Reservations -----\n")

	for index, user := range *bookings {
		fmt.Printf("Reservation %v:\nName: %v    Tickets: %v\n", index+1, user.firstName, user.tickets)
	}
}

func sendTicket(user Booking) {
	time.Sleep(10 * time.Second)

	var ticket = fmt.Sprintf("%v tickets for %v %v", user.tickets, user.firstName, user.lastName)

	fmt.Printf("##########\n")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, user.email)
	fmt.Printf("##########\n\n")

	WG.Done()
}
