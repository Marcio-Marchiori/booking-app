package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v mate!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get the tickets here.")

	var userName string
	var userTickets int
	// ask user for input
	fmt.Scan(&userName)

	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)
}
