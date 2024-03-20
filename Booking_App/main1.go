package main

import "fmt"

func main() {
	var conferenceName string = "Go Conference" //var declaration
	const conferenceTickets int = 50
	//it is declared const not var so it will show the error
	//if we write var conferenceTickets then its value can be changed again and again

	var remainingTickets int = 50
	//now if we use var remainingTickets uint = 50
	//then we can not set the value of remainingTickets to -ve
	//but if it is int instead of uint then no issue it will cover all the values in the integer range

	// Println is used for the output to be print on the new line
	// \n can also be used for the new line

	fmt.Printf("conferenceTickets is %T, ramainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceName)
	//%T gives the type of the variable it is

	//fmt.Println("Welcome to", conferenceName, "booking application") -> we can write this too
	fmt.Printf("\nWelcome to %v booking application\n", conferenceName)

	// fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets, "are still available.")
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)

	fmt.Println("Get your tickets here to attend the conference")

	// for user input

	// ask user for their name
	var firstName string
	var userTickets int
	// firstName = "Tom"
	// for user input
	//fmt.Scan(firstName)
	//so this will not wait for the user input
	// so we will do this --> pointer
	fmt.Println("Enter yout first name")
	fmt.Scan(&firstName)
	// this will be used to take user input

	userTickets = 2
	fmt.Printf("User %v bookd %v tickets.\n", firstName, userTickets)

}
