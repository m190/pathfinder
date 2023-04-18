package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetFilePath(t *testing.T) {
	tests := []struct {
		flights []Flight
		path    []string
	}{
		{
			flights: []Flight{
				{Source: "SFO", Destination: "EWR"},
			},
			path: []string{"SFO", "EWR"},
		},
		{
			flights: []Flight{
				{Source: "ATL", Destination: "EWR"},
				{Source: "SFO", Destination: "ATL"},
			},
			path: []string{"SFO", "EWR"},
		},
		{
			flights: []Flight{
				{Source: "IND", Destination: "EWR"},
				{Source: "SFO", Destination: "ATL"},
				{Source: "GSO", Destination: "IND"},
				{Source: "ATL", Destination: "GSO"},
			},
			path: []string{"SFO", "EWR"},
		},
		{
			flights: []Flight{
				{Source: "A", Destination: "B"},
				{Source: "B", Destination: "C"},
				{Source: "C", Destination: "D"},
				{Source: "D", Destination: "E"},
				{Source: "E", Destination: "B"},
			},
			path: []string{"A", "B"},
		},
		{
			flights: []Flight{
				{Source: "A", Destination: "B"},
				{Source: "B", Destination: "C"},
				{Source: "C", Destination: "D"},
				{Source: "D", Destination: "C2"},
				{Source: "C2", Destination: "D"},
				{Source: "D", Destination: "E"},
			},
			path: []string{"A", "E"},
		},
		{
			flights: []Flight{
				{Source: "A", Destination: "B"},
				{Source: "B", Destination: "C"},
				{Source: "C", Destination: "D"},
				{Source: "D", Destination: "C2"},
				{Source: "C2", Destination: "B2"},
				{Source: "B2", Destination: "C"},
				{Source: "C", Destination: "D"},
			},
			path: []string{"A", "D"},
		},
	}

	for _, tt := range tests {
		t.Run(strings.Join(tt.path, "->"), func(t *testing.T) {
			path := getFlightPath(tt.flights)
			if !reflect.DeepEqual(path, tt.path) {
				t.Errorf("expected %v, actual %v", tt.path, path)
			}
		})
	}
}
