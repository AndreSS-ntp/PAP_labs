package repository

import (
	"encoding/json"
	"fmt"
	"github.com/AndreSS-ntp/PAP_labs/tree/main/lab6/internal/domain"
	"os"
)

func LoadFlightsFromFile(filename string) error {
	domain.Mu.Lock()
	defer domain.Mu.Unlock()

	file, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("could not read file: %v", err)
	}

	var flightsList []domain.Flight
	if err := json.Unmarshal(file, &flightsList); err != nil {
		return fmt.Errorf("could not unmarshal flights data: %v", err)
	}

	for _, flight := range flightsList {
		domain.Flights[flight.ID] = flight
	}

	return nil
}
