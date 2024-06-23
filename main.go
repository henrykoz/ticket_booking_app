package main

import (
	"fmt"
	"hello/helper"
	"time"
)

const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
//empty list of maps
var bookings = make([]UserData, 0)

type UserData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
} 

func main() {

	greetUsers()

	for  {
		//user input
		firstName, lastName, email, userTickets := getUserInput()

		// validation
		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			
			// booking
			bookTicket(userTickets, firstName , lastName, email)
			sendTicket(userTickets, firstName , lastName, email)

			// call function print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)


			if remainingTickets == 0 {
			// end program
			fmt.Println("Our conference is fully booked.")
			break
			}
		} else {
			if !isValidName {
				fmt.Println("name is too short")
			}
			if !isValidEmail {
				fmt.Println("wrong email address")
			}
			if !isValidTicketNumber {
				fmt.Println("number of tickets is invalid")
			}
			
		}

		
	}

}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("Total tickets available: %v Remaining tickets: %v \n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets now!")
}

func getFirstNames() []string {
	firstNames := []string{}
			for _, booking := range bookings {
				firstNames = append(firstNames, booking.firstName)
			}
			fmt.Printf("These are all our bookings: %v\n", bookings)
			return firstNames
			
}


func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user their name
	fmt.Println("Enter your first name:")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name:")
	fmt.Scan(&lastName)

	fmt.Println("Enter your e-mail")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets:")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//create a map for user
	var userData = UserData {
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation e-mail at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("##########")
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println("##########")
}