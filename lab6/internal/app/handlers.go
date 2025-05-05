package app

import (
	"encoding/json"
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab6/internal/domain"
	"net/http"
)

// Получить все рейсы
func GetFlightsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	domain.Mu.Lock()
	defer domain.Mu.Unlock()

	flightsList := make([]domain.Flight, 0, len(domain.Flights))
	for _, flight := range domain.Flights {
		flightsList = append(flightsList, flight)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(flightsList)
}

// Получить рейсы по маршруту
func GetFlightsByRouteHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	origin := r.URL.Query().Get("origin")
	destination := r.URL.Query().Get("destination")

	if origin == "" || destination == "" {
		http.Error(w, "Origin and destination parameters are required", http.StatusBadRequest)
		return
	}

	domain.Mu.Lock()
	defer domain.Mu.Unlock()

	var result []domain.Flight
	for _, flight := range domain.Flights {
		if flight.Origin == origin && flight.Destination == destination {
			result = append(result, flight)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// Получить количество свободных мест на рейсе
func GetFlightSeatsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	flightID := r.URL.Query().Get("flight_id")
	if flightID == "" {
		http.Error(w, "Flight ID is required", http.StatusBadRequest)
		return
	}

	domain.Mu.Lock()
	defer domain.Mu.Unlock()

	flight, exists := domain.Flights[flightID]
	if !exists {
		http.Error(w, "Flight not found", http.StatusNotFound)
		return
	}

	response := map[string]interface{}{
		"flight_id": flightID,
		"seats":     flight.Seats,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Забронировать билет
func BookTicketHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		FlightID  string `json:"flight_id"`
		Passenger string `json:"passenger"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	domain.Mu.Lock()
	defer domain.Mu.Unlock()

	flight, exists := domain.Flights[request.FlightID]
	if !exists {
		http.Error(w, "Flight not found", http.StatusNotFound)
		return
	}

	if flight.Seats <= 0 {
		http.Error(w, "No seats available", http.StatusConflict)
		return
	}

	// Уменьшаем количество свободных мест
	flight.Seats--
	domain.Flights[request.FlightID] = flight

	// Создаем билет
	ticketID := fmt.Sprintf("ticket-%d", len(domain.Tickets)+1)
	domain.Tickets[ticketID] = domain.Ticket{
		FlightID:  request.FlightID,
		Passenger: request.Passenger,
	}

	response := map[string]interface{}{
		"ticket_id": ticketID,
		"flight_id": request.FlightID,
		"passenger": request.Passenger,
		"message":   "Ticket booked successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
