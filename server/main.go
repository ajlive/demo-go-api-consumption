package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const banjoID = "05cc936a-6068-11ed-8427-32bc7acc9df8"

const banjoPayload = `
{
	"id": "%v",
	"name": "Banjo",
	"species": "cat",
	"color": "orange",
	"age": 5,
	"weight": 15.5,
	"location": {
		"city": "Denver",
		"state": "CO"
	}
}
`

const errorPayload = `
{
	"error": "%v"
}
`

func run() error {
	writeSuccess := func(w http.ResponseWriter, payload string) {
		w.Header().Add("Content-Type", "application/json")
		payload = strings.TrimSpace(payload)
		w.Write([]byte(payload))
	}

	writeError := func(w http.ResponseWriter, status int, err error) {
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		payload := fmt.Sprintf(errorPayload, err)
		payload = strings.TrimSpace(payload)
		w.Write([]byte(payload))
	}

	http.HandleFunc("/api/pets", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()
		ids, ok := query["id"]
		if !ok || len(ids) != 1 {
			writeError(w, 422, fmt.Errorf("url must have one id query parameter"))
			return
		}
		id := ids[0]

		if id != banjoID {
			writeError(w, 422, fmt.Errorf("no pet with id %v", id))
			return
		}

		payload := fmt.Sprintf(banjoPayload, id)
		writeSuccess(w, payload)
	})

	fmt.Printf("starting pet server on :8989\nexample URL:\nhttp://localhost:8989/api/pets?id=%v\n", banjoID)
	return http.ListenAndServe(":8989", nil)
}

func main() {
	log.Fatal(run())
}
