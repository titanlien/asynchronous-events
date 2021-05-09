package main

import "net/http"
import "encoding/json"

type Coaster struct {
	Name       string `json:"name"`
	Manufactor string `json:"manufactor"`
	ID         string `json:"id"`
	InPark     string `json:"inPark"`
	Height     int    `json:"height"`
}

type coasterHandlers struct {
	store map[string]Coaster
}

func (h *coasterHandlers) get(w http.ResponseWriter, r *http.Request) {
	coasters := make([]Coaster, len(h.store))
	i := 0
	for _, coaster := range h.store {
		coasters[i] = coaster
		i++
	}

	jsonBytes, err := json.Marshal(coasters)
	if err != nil {
		// TODO
	}
	w.Write(jsonBytes)
}

func newCoasterHandlers() *coasterHandlers {
	return &coasterHandlers{
		store: map[string]Coaster{
			"id1": Coaster{
				Name:       "Fury 325",
				Height:     99,
				ID:         "id1",
				InPark:     "Carowinds",
				Manufactor: "B+M",
			},
		},
	}
}

func main() {
	coasterHandlers := newCoasterHandlers()
	http.HandleFunc("/coasters", coasterHandlers.get)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
