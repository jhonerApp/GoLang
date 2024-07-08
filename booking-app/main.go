package main

import (
	"fmt"
	"strings"
)

func main() {

	// using := , you dont need to use const or var in your variable

	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 30
	bookings := []string{} //or var bookings []string or var bookings = []string{} -> syntax to assign value in array variable

	fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n ", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("We have total of %v, conferenceTickets, tickets and %v, are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets")

	//For loop

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//& = pointers or special variable / convert 0x000134
		fmt.Println(("Enter your first name: "))
		fmt.Scan(&firstName)

		fmt.Println(("Enter your last name: "))
		fmt.Scan(&lastName)

		fmt.Println(("Enter your email: "))
		fmt.Scan(&email)

		fmt.Println(("Enter number of  tickets: "))
		fmt.Scan(&userTickets)

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole array: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for  %v\n", remainingTickets, conferenceName)

		// range
		//> range interates over the elements for different data structures(so not only arrays and slices)
		//> for arrays and slices, range provdes the index and value for each element

		// strings.Fields()
		//> splits the strings with white space as separator
		//> And returns a slice with the split elements
		//> Ex: Nicole Smith ==> ["Nicole","Smith"]

		//Blank identifier
		//> _ -> to ignore a variable you dont want to use
		//> So with Go you need to make unused variables explicit

		firstNames := []string{}
		for _, booking := range bookings {

			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings are: %v\n", firstNames)
	}

}
