package main

import "fmt"

// Simple individual tasks
func makeHotelReservation() {
	fmt.Println("Done making hotel reservation")
}

func bookFlightTickets() {
	fmt.Println("Done booking flight tickets")
}

func orderADress() {
	fmt.Println("Done ordering a dress")
}

func payCreditDebitCardBills() {
	fmt.Println("Done paying Credit Card bills")
}

//Tasks that will be executed in parts
// Writing Mail
func writeAMail() {
	fmt.Println("Wrote 1/3rd of the mail.")
	go continueWritingMail1()
}

func continueWritingMail1() {
	fmt.Println("Wrote 2/3rds of the mail.")
	go continueWritingMail2()
}

func continueWritingMail2() {
	fmt.Println("Done writing the mail.")
}

// Listening to Audio Book
func listenToAudioBook() {
	fmt.Println("Listened to 10 minutes of audio book.")
	go continueListeningToAudioBook()
}
func continueListeningToAudioBook() {
	fmt.Println("Done listening to audio book.")
}

//All the tasks we want to complete in the day. Not that we do not include the sub-tasks here
var listOfTasks = []func(){
	makeHotelReservation, bookFlightTickets, orderADress, payCreditDebitCardBills, writeAMail, listenToAudioBook,
}

func main() {
	//1. Linear serial task execution
	for _, task := range listOfTasks {
		task()
	}
}
