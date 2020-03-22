package main

import (
	"fmt"
	"time"
)

// Flights data such as arrival and departures time
type Flights interface {
	// returns arrival time
	arrival() time.Time
	// returns departures time
	departures() time.Time
}

// Ryanair is a test structure which contains just flight number field
type Ryanair struct {
	// flight number
	number string
}

// Norwegian is a test structure which contains just flight number field
type Norwegian struct {
	// flight number
	number string
}

// interface functions implementation for Ryanair structure
func (r Ryanair) arrival() time.Time {
	// arrival should be tomorrow + 2h
	return time.Now().Add(+26 * time.Hour)
}

func (r Ryanair) departures() time.Time {
	// departure should be tommorow
	return time.Now().Add(+24 * time.Hour)
}

// interface functions implementation for Norwegian structure
func (n Norwegian) arrival() time.Time {
	// arrival should be in 1h
	return time.Now().Add(+1 * time.Hour)
}

func (n Norwegian) departures() time.Time {
	// flight has been departed
	return time.Now().Add(-1 * time.Hour)
}

// Common function for both Norwegian and Ryanair
func status(flight Flights) string {
	t := time.Now()
	switch {
	case t.Before(flight.departures()):
		return "arrives later"
	case t.After(flight.arrival()):
		return "has been departed"
	default:
		return "not yet departed"
	}
}

func main() {
	var ryanair = Ryanair{
		number: "RR111",
	}

	var norwegian = Norwegian{
		number: "NN111",
	}

	fmt.Printf("Flight %v %v \n", ryanair.number, status(ryanair))
	fmt.Printf("Flight %v %v \n", norwegian.number, status(norwegian))
}
