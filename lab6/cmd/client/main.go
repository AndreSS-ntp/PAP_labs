package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Flight Booking System Client")
	fmt.Println("===========================")

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. View all flights")
		fmt.Println("2. Search flights by route")
		fmt.Println("3. Check available seats")
		fmt.Println("4. Book a ticket")
		fmt.Println("5. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			getAllFlights()
		case 2:
			searchFlightsByRoute()
		case 3:
			checkAvailableSeats()
		case 4:
			bookTicket()
		case 5:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, please try again")
		}
	}
}

func getAllFlights() {
	resp, err := http.Get("http://localhost:8080/flights")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var flights []struct {
		ID          string `json:"id"`
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
		Date        string `json:"date"`
		Seats       int    `json:"seats"`
	}

	if err := json.Unmarshal(body, &flights); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("\nAll Available Flights:")
	for _, flight := range flights {
		fmt.Printf("ID: %s, %s -> %s, Date: %s, Seats: %d\n",
			flight.ID, flight.Origin, flight.Destination, flight.Date, flight.Seats)
	}
}

func searchFlightsByRoute() {
	var origin, destination string
	fmt.Print("Enter origin: ")
	fmt.Scanln(&origin)
	fmt.Print("Enter destination: ")
	fmt.Scanln(&destination)

	url := fmt.Sprintf("http://localhost:8080/flights/route?origin=%s&destination=%s", origin, destination)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var flights []struct {
		ID          string `json:"id"`
		Origin      string `json:"origin"`
		Destination string `json:"destination"`
		Date        string `json:"date"`
		Seats       int    `json:"seats"`
	}

	if err := json.Unmarshal(body, &flights); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\nFlights from %s to %s:\n", origin, destination)
	for _, flight := range flights {
		fmt.Printf("ID: %s, Date: %s, Seats: %d\n",
			flight.ID, flight.Date, flight.Seats)
	}
}

func checkAvailableSeats() {
	var flightID string
	fmt.Print("Enter flight ID: ")
	fmt.Scanln(&flightID)

	url := fmt.Sprintf("http://localhost:8080/flights/seats?flight_id=%s", flightID)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var result struct {
		FlightID string `json:"flight_id"`
		Seats    int    `json:"seats"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\nFlight %s has %d available seats\n", result.FlightID, result.Seats)
}

func bookTicket() {
	var flightID, passenger string
	fmt.Print("Enter flight ID: ")
	fmt.Scanln(&flightID)
	fmt.Print("Enter passenger name: ")
	fmt.Scanln(&passenger)

	requestBody, err := json.Marshal(map[string]string{
		"flight_id": flightID,
		"passenger": passenger,
	})
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	resp, err := http.Post("http://localhost:8080/tickets/book", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var result struct {
		TicketID string `json:"ticket_id"`
		Message  string `json:"message"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\n%s. Ticket ID: %s\n", result.Message, result.TicketID)
}
