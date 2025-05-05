package main

import (
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab6/internal/app"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab6/internal/repository"
	"log"
	"net/http"
)

func main() {
	err := repository.LoadFlightsFromFile("flights.json")
	if err != nil {
		log.Fatalf("Failed to load flights: %v", err)
	}

	http.HandleFunc("/flights", app.GetFlightsHandler)
	http.HandleFunc("/flights/route", app.GetFlightsByRouteHandler)
	http.HandleFunc("/flights/seats", app.GetFlightSeatsHandler)
	http.HandleFunc("/tickets/book", app.BookTicketHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
