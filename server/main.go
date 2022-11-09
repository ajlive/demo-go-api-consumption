package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const banjoID = "05cc936a-6068-11ed-8427-32bc7acc9df8"

type errorResponse struct {
	Error string `json:"error"`
}

type location struct {
	City  string `json:"city"`
	State string `json:"state"`
}
type pet struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Species  string   `json:"species"`
	Color    string   `json:"color"`
	Age      int64    `json:"age"`
	Weight   float64  `json:"weight"`
	Location location `json:"location"`
}

var petsRepository = map[string]pet{
	banjoID: {
		ID:      banjoID,
		Name:    "Banjo",
		Species: "cat",
		Color:   "orange",
		Age:     5,
		Weight:  15.5,
		Location: location{
			City:  "Denver",
			State: "CO",
		},
	},
}

func run() error {
	writeSuccess := func(w http.ResponseWriter, resp any) {
		w.Header().Add("Content-Type", "application/json")
		payload, err := json.MarshalIndent(resp, "", "\t")
		if err != nil {
			panic(err) // panic, like throwing an exception except designed to crash the app; used for programmer errors
			// can be used with built-in recover() function, but expected errors are better handled with error returns
		}
		w.Write(payload)
	}

	writeError := func(w http.ResponseWriter, status int, err error) {
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		resp := errorResponse{
			Error: err.Error(),
		}
		payload, err := json.MarshalIndent(resp, "", "\t")
		if err != nil {
			panic(err)
		}
		w.Write(payload)
	}

	http.HandleFunc("/api/pets", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query() // map of query parameters map[string][]string
		ids, ok := query["id"] // second return, "ok" bool, says whether the key was found in the map
		if !ok || len(ids) != 1 {
			writeError(w, 422, fmt.Errorf("url must have one id query parameter"))
			return
		}
		id := ids[0]

		if id != banjoID {
			writeError(w, 422, fmt.Errorf("no pet with id %v", id))
			return
		}

		banjo := petsRepository[id]
		writeSuccess(w, banjo)
	})

	fmt.Printf("starting pet server on :8989\nexample URL:\nhttp://localhost:8989/api/pets?id=%v\n", banjoID)
	return http.ListenAndServe(":8989", nil)
}

func main() {
	log.Fatal(run())
}
