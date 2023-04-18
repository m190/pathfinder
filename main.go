package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Flight struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
}

func main() {
	http.HandleFunc("/calculate", flightPathHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func flightPathHandler(w http.ResponseWriter, r *http.Request) {
	var flights [][]string
	err := json.NewDecoder(r.Body).Decode(&flights)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var convertedFlights []Flight
	for _, f := range flights {
		if len(f) != 2 {
			http.Error(w, "Invalid flight data", http.StatusBadRequest)
			return
		}
		convertedFlights = append(convertedFlights, Flight{Source: f[0], Destination: f[1]})
	}

	path := getFlightPath(convertedFlights)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(path)
}

func getFlightPath(flights []Flight) []string {
	path := make([]string, 0)
	sources := make(map[string]struct{})
	destinations := make(map[string]struct{})
	for _, flight := range flights {
		sources[flight.Source] = struct{}{}
		destinations[flight.Destination] = struct{}{}
	}

	var start, end string
	for source := range sources {
		if _, ok := destinations[source]; !ok {
			start = source
			break
		}
	}
	for destination := range destinations {
		if _, ok := sources[destination]; !ok {
			end = destination
			break
		}
	}
	if start != "" && end != "" {
		return []string{start, end}
	}

	path = append(path, start)
outer:
	for {
		for i, flight := range flights {
			if flight.Source == start {
				path = append(path, flight.Destination)
				start = flight.Destination
				flights = append(flights[:i], flights[i+1:]...)
				continue outer
			}
		}
		return []string{path[0], path[len(path)-1]}
	}
}
