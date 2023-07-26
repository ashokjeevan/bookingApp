package main
import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets = 50
	var remainingTickets uint = 50

	var firstName string
	var lastName string
	var email string
	var noOfTickets uint

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and we have %v tickets available\n", conferenceTickets, remainingTickets)

	// taking user input
	fmt.Println("Please enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Please enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Please enter your email address")
	fmt.Scan(&email)

	fmt.Println("Please enter how many tickets you want")
	fmt.Scan(&noOfTickets)

	fmt.Printf("Hi %v %v, thank you for booking %v tickets for %v\n", firstName, lastName, noOfTickets, conferenceName)
	fmt.Printf("Your tickets will be sent to %v\n", email)

	remainingTickets = remainingTickets - noOfTickets

	fmt.Printf("The %v has %v tickets remaining\n", conferenceName, remainingTickets)

}