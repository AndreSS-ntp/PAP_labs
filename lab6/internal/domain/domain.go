package domain

import "sync"

type Flight struct {
	ID          string `json:"id"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
	Seats       int    `json:"seats"`
}

type Ticket struct {
	FlightID  string `json:"flight_id"`
	Passenger string `json:"passenger"`
}

var (
	Flights = make(map[string]Flight)
	Tickets = make(map[string]Ticket)
	Mu      sync.Mutex
)
