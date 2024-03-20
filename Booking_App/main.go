package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference" //var declaration
	const conferenceTickets int = 50
	var remainingTickets int = 50

	fmt.Printf("\nconferenceTickets is %T, ramainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("\nWelcome to %v booking application\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	fmt.Println("Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email")
	fmt.Scan(&email)

	fmt.Println("Enter you first name")
	fmt.Scan(&firstName)

	fmt.Printf("User %v book %v tickets.\n", firstName, userTickets)
}
