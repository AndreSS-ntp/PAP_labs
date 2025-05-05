package main

import (
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab6/internal/app"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab6/internal/domain"
	"log"
	"net/http"
)

func main() {
	domain.Flights["1"] = domain.Flight{ID: "1", Origin: "Moscow", Destination: "Saint-Petersburg", Date: "2023-12-01", Seats: 50}
	domain.Flights["2"] = domain.Flight{ID: "2", Origin: "Moscow", Destination: "Sochi", Date: "2023-12-02", Seats: 30}
	domain.Flights["3"] = domain.Flight{ID: "3", Origin: "Saint-Petersburg", Destination: "Sochi", Date: "2023-12-03", Seats: 40}

	http.HandleFunc("/flights", app.GetFlightsHandler)
	http.HandleFunc("/flights/route", app.GetFlightsByRouteHandler)
	http.HandleFunc("/flights/seats", app.GetFlightSeatsHandler)
	http.HandleFunc("/tickets/book", app.BookTicketHandler)

	fmt.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
