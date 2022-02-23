package main

import "fmt"

func main() {

	var conference = "GolangCon"
	const conferenceTickets = 50
	var remainingTickets = 0

	fmt.Printf("Welcome to %v! Booking Application\n", conference)
	fmt.Printf("There are %v tickets available for sale and %v tickets remaining\n", conferenceTickets, remainingTickets)

	var bookings [50]string
	//ask user for their name
	var name string
	var userTickets int
	fmt.Print("Please enter your name:\n ")
	fmt.Scan(&name)
	fmt.Printf("Please enter the number of tickets you would like to purchase:\n")
	fmt.Scan(&userTickets)
	remainingTickets = conferenceTickets - userTickets
	bookings[0] = name

	fmt.Printf("%v, you have purchased %v tickets. Now the remaining tickets are %v \n", name, userTickets, remainingTickets)
	fmt.Printf("The bookings are: %v \n", bookings)

}
